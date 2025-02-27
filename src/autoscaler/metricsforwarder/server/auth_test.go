package server_test

import (
	"autoscaler/fakes"
	"autoscaler/metricsforwarder/config"
	"autoscaler/metricsforwarder/server/auth"
	"autoscaler/models"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"path"
	"time"

	"code.cloudfoundry.org/lager"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"

	"github.com/patrickmn/go-cache"
)

var _ = Describe("Authentication", func() {

	var (
		authTest                   *auth.Auth
		credentialCache            cache.Cache
		policyDB                   *fakes.FakePolicyDB
		resp                       *httptest.ResponseRecorder
		req                        *http.Request
		body                       []byte
		vars                       map[string]string
		metricsForwarderMtlsConfig config.MetricsForwarderConfig
	)

	BeforeEach(func() {
		logger := lager.NewLogger("auth-test")
		policyDB = &fakes.FakePolicyDB{}
		credentialCache = *cache.New(10*time.Minute, -1)
		vars = make(map[string]string)
		resp = httptest.NewRecorder()
		metricsForwarderMtlsConfig = config.MetricsForwarderConfig{
			TLS: models.TLSCerts{
				CACertFile: path.Join("..", "..", "..", "..", "test-certs", "valid-mtls-local-ca-combined.crt"),
			}}
		authTest = auth.New(logger, policyDB, credentialCache, 10*time.Minute, metricsForwarderMtlsConfig.TLS.CACertFile)
		credentialCache.Flush()
	})

	Describe("Basic Auth tests for publish metrics endpoint", func() {

		credentials := &models.Credential{
			Username: "$2a$10$YnQNQYcvl/Q2BKtThOKFZ.KB0nTIZwhKr5q1pWTTwC/PUAHsbcpFu",
			Password: "$2a$10$6nZ73cm7IV26wxRnmm5E1.nbk9G.0a4MrbzBFPChkm5fPftsUwj9G",
		}

		When("a valid request to publish custom metrics comes", func() {
			When("credentials exists in the cache", func() {
				It("should get the credentials from cache without searching from database and calls next handler", func() {
					credentialCache.Set("an-app-id", credentials, 10*time.Minute)
					req = CreateRequest(body)
					req.Header.Add("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
					vars["appid"] = "an-app-id"
					nextCalled := 0
					nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						nextCalled = nextCalled + 1
					})

					authTest.AuthenticateHandler(nextFunc)(resp, req, vars)
					Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
					Expect(resp.Code).To(Equal(http.StatusOK))
					Expect(nextCalled).To(Equal(1))
				})

			})

			When("credentials do not exists in the cache but exist in the database", func() {
				It("should: get the credentials from database, add it to the cache and calls next handler", func() {
					req = CreateRequest(body)
					req.Header.Add("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
					vars["appid"] = "an-app-id"
					nextCalled := 0
					nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						nextCalled = nextCalled + 1
					})
					policyDB.GetCredentialReturns(credentials, nil)

					authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

					Expect(policyDB.GetCredentialCallCount()).To(Equal(1))
					Expect(resp.Code).To(Equal(http.StatusOK))
					Expect(nextCalled).To(Equal(1))
					//fills the cache
					_, found := credentialCache.Get("an-app-id")
					Expect(found).To(Equal(true))
				})

			})

			Context("when credentials neither exists in the cache nor exist in the database", func() {
				It("should search in both cache & database and returns status code 401", func() {
					req = CreateRequest(body)
					req.Header.Add("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
					vars["appid"] = "an-app-id"
					nextCalled := 0
					nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						nextCalled = nextCalled + 1
					})
					policyDB.GetCredentialReturns(nil, sql.ErrNoRows)

					authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

					Expect(policyDB.GetCredentialCallCount()).To(Equal(1))
					Expect(resp.Code).To(Equal(http.StatusUnauthorized))
					errJson := &models.ErrorResponse{}
					err := json.Unmarshal(resp.Body.Bytes(), errJson)
					Expect(err).ToNot(HaveOccurred())
					Expect(errJson).To(Equal(&models.ErrorResponse{
						Code:    "Unauthorized",
						Message: "Unauthorized",
					}))
					Expect(nextCalled).To(Equal(0))
				})

			})

			Context("when a stale credentials exists in the cache", func() {
				It("should search in the database and calls next handler", func() {
					req = CreateRequest(body)
					credentialCache.Set("an-app-id", &models.Credential{Username: "some-stale-hashed-username", Password: "some-stale-hashed-password"}, 10*time.Minute)
					policyDB.GetCredentialReturns(credentials, nil)
					req.Header.Add("Authorization", "Basic dXNlcm5hbWU6cGFzc3dvcmQ=")
					vars["appid"] = "an-app-id"
					nextCalled := 0
					nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						nextCalled = nextCalled + 1
					})

					authTest.AuthenticateHandler(nextFunc)(resp, req, vars)
					Expect(policyDB.GetCredentialCallCount()).To(Equal(1))
					Expect(resp.Code).To(Equal(http.StatusOK))
					Expect(nextCalled).To(Equal(1))
				})
			})
		})
	})

	Describe("MTLS Auth tests for publish metrics endpoint", func() {
		const validClientCert1 = "../../../../test-certs/validmtls_client-1.crt"
		When("correct xfcc header with correct CA is supplied for cert 1", func() {
			It("should call next handler", func() {
				req = CreateRequest(body)
				req.Header.Add("X-Forwarded-Client-Cert", MustReadXFCCcert(validClientCert1))
				vars["appid"] = "an-app-id"
				nextCalled := 0
				nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					nextCalled = nextCalled + 1
				})

				authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

				Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
				Expect(resp.Code).To(Equal(http.StatusOK))
				Expect(nextCalled).To(Equal(1))
			})
		})

		When("correct xfcc header with correct CA is supplied for cert 2", func() {
			It("should call next handler", func() {
				req = CreateRequest(body)
				const validClientCert2 = "../../../../test-certs/validmtls_client-2.crt"
				req.Header.Add("X-Forwarded-Client-Cert", MustReadXFCCcert(validClientCert2))
				vars["appid"] = "an-app-id"
				nextCalled := 0
				nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					nextCalled = nextCalled + 1
				})

				authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

				Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
				Expect(resp.Code).To(Equal(http.StatusOK))
				Expect(nextCalled).To(Equal(1))
			})
		})

		When("correct xfcc header including \"'s around the cert", func() {
			It("should call next handler", func() {
				req = CreateRequest(body)
				req.Header.Add("X-Forwarded-Client-Cert", fmt.Sprintf("%q", MustReadXFCCcert(validClientCert1)))
				vars["appid"] = "an-app-id"
				nextCalled := 0
				nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					nextCalled = nextCalled + 1
				})

				authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

				Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
				Expect(resp.Code).To(Equal(http.StatusOK))
				Expect(nextCalled).To(Equal(1))
			})
		})

		When("correct xfcc header with invalid CA is supplied", func() {
			It("should return status code 401", func() {
				req = CreateRequest(body)
				req.Header.Add("X-Forwarded-Client-Cert", MustReadXFCCcert("../../../../test-certs/invalidmtls_client.crt"))
				vars["appid"] = "an-app-id"
				nextCalled := 0
				nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					nextCalled = nextCalled + 1
				})

				authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

				Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
				Expect(resp.Body.String()).To(Equal(`{"code":"Unauthorized","message":"Unauthorized"}`))
				Expect(nextCalled).To(Equal(0))
			})
		})

		When("correct xfcc header with no CA is supplied", func() {
			It("should return status code 401", func() {
				req = CreateRequest(body)
				req.Header.Add("X-Forwarded-Client-Cert", MustReadXFCCcert("../../../../test-certs/nosignmtls_client.crt"))
				vars["appid"] = "an-app-id"
				nextCalled := 0
				nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					nextCalled = nextCalled + 1
				})

				authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

				Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
				Expect(resp.Body.String()).To(Equal(`{"code":"Unauthorized","message":"Unauthorized"}`))
				Expect(nextCalled).To(Equal(0))
			})
		})

		When("valid cert with wrong app-id is supplied", func() {
			It("should return status code 403", func() {
				req = CreateRequest(body)
				req.Header.Add("X-Forwarded-Client-Cert", MustReadXFCCcert(validClientCert1))
				nextCalled := 0
				nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					nextCalled = nextCalled + 1
				})

				vars["appid"] = "wrong-an-app-id"
				authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

				Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
				Expect(resp.Code).To(Equal(http.StatusForbidden))
				Expect(resp.Body.String()).To(Equal(`{"code":"Forbidden","message":"Unauthorized"}`))
				Expect(nextCalled).To(Equal(0))
			})
		})

		When("expired cert with correct app-id is supplied", func() {
			It("should return status code 401", func() {
				req = CreateRequest(body)
				req.Header.Add("X-Forwarded-Client-Cert", MustReadXFCCcert("../../../../test-certs/expiredmtls_client.crt"))
				nextCalled := 0
				nextFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					nextCalled = nextCalled + 1
				})

				vars["appid"] = "an-app-id"
				authTest.AuthenticateHandler(nextFunc)(resp, req, vars)

				Expect(policyDB.GetCredentialCallCount()).To(Equal(0))
				Expect(resp.Code).To(Equal(http.StatusUnauthorized))
				Expect(nextCalled).To(Equal(0))
			})
		})
	})

})

func MustReadXFCCcert(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	Expect(err).ShouldNot(HaveOccurred())
	block, _ := pem.Decode(file)
	Expect(block).ShouldNot(BeNil())
	return base64.StdEncoding.EncodeToString(block.Bytes)
}
