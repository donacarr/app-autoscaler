package config

import (
	"autoscaler/cf"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"

	"autoscaler/db"
	"autoscaler/helpers"
	"autoscaler/models"
)

const (
	DefaultLoggingLevel  = "info"
	DefaultMaxAmount     = 10
	DefaultValidDuration = 1 * time.Second
)

type ServerConfig struct {
	Port int             `yaml:"port"`
	TLS  models.TLSCerts `yaml:"tls"`
}

var defaultBrokerServerConfig = ServerConfig{
	Port: 8080,
}

var defaultPublicApiServerConfig = ServerConfig{
	Port: 8081,
}

var defaultLoggingConfig = helpers.LoggingConfig{
	Level: DefaultLoggingLevel,
}

type DBConfig struct {
	BindingDB db.DatabaseConfig `yaml:"binding_db"`
	PolicyDB  db.DatabaseConfig `yaml:"policy_db"`
}

type SchedulerConfig struct {
	SchedulerURL   string          `yaml:"scheduler_url"`
	TLSClientCerts models.TLSCerts `yaml:"tls"`
}
type ScalingEngineConfig struct {
	ScalingEngineUrl string          `yaml:"scaling_engine_url"`
	TLSClientCerts   models.TLSCerts `yaml:"tls"`
}

type MetricsCollectorConfig struct {
	MetricsCollectorUrl string          `yaml:"metrics_collector_url"`
	TLSClientCerts      models.TLSCerts `yaml:"tls"`
}

type EventGeneratorConfig struct {
	EventGeneratorUrl string          `yaml:"event_generator_url"`
	TLSClientCerts    models.TLSCerts `yaml:"tls"`
}
type MetricsForwarderConfig struct {
	MetricsForwarderUrl     string `yaml:"metrics_forwarder_url"`
	MetricsForwarderMtlsUrl string `yaml:"metrics_forwarder_mtls_url"`
}

type QuotaManagementConfig struct {
	API               string `yaml:"api"`
	ClientID          string `yaml:"client_id"`
	Secret            string `yaml:"secret"`
	TokenURL          string `yaml:"oauth_url"`
	SkipSSLValidation bool   `yaml:"skip_ssl_validation"`
}

type PlanDefinition struct {
	PlanCheckEnabled  bool `yaml:"planCheckEnabled"`
	SchedulesCount    int  `yaml:"schedules_count"`
	ScalingRulesCount int  `yaml:"scaling_rules_count"`
	PlanUpdateable    bool `yaml:"plan_updateable"`
}

type BrokerCredentialsConfig struct {
	BrokerUsername     string `yaml:"broker_username"`
	BrokerUsernameHash []byte `yaml:"broker_username_hash"`
	BrokerPassword     string `yaml:"broker_password"`
	BrokerPasswordHash []byte `yaml:"broker_password_hash"`
}

type Config struct {
	Logging              helpers.LoggingConfig     `yaml:"logging"`
	BrokerServer         ServerConfig              `yaml:"broker_server"`
	PublicApiServer      ServerConfig              `yaml:"public_api_server"`
	DB                   DBConfig                  `yaml:"db"`
	BrokerCredentials    []BrokerCredentialsConfig `yaml:"broker_credentials"`
	APIClientId          string                    `yaml:"api_client_id"`
	QuotaManagement      *QuotaManagementConfig    `yaml:"quota_management"`
	PlanCheck            *PlanCheckConfig          `yaml:"plan_check"`
	CatalogPath          string                    `yaml:"catalog_path"`
	CatalogSchemaPath    string                    `yaml:"catalog_schema_path"`
	DashboardRedirectURI string                    `yaml:"dashboard_redirect_uri"`
	PolicySchemaPath     string                    `yaml:"policy_schema_path"`
	Scheduler            SchedulerConfig           `yaml:"scheduler"`
	ScalingEngine        ScalingEngineConfig       `yaml:"scaling_engine"`
	MetricsCollector     MetricsCollectorConfig    `yaml:"metrics_collector"`
	EventGenerator       EventGeneratorConfig      `yaml:"event_generator"`
	CF                   cf.CFConfig               `yaml:"cf"`
	UseBuildInMode       bool                      `yaml:"use_buildin_mode"`
	InfoFilePath         string                    `yaml:"info_file_path"`
	MetricsForwarder     MetricsForwarderConfig    `yaml:"metrics_forwarder"`
	Health               models.HealthConfig       `yaml:"health"`
	RateLimit            models.RateLimitConfig    `yaml:"rate_limit"`
}

type PlanCheckConfig struct {
	PlanDefinitions map[string]PlanDefinition `yaml:"plan_definitions"`
}

func LoadConfig(reader io.Reader) (*Config, error) {
	conf := &Config{
		Logging:         defaultLoggingConfig,
		BrokerServer:    defaultBrokerServerConfig,
		PublicApiServer: defaultPublicApiServerConfig,
		UseBuildInMode:  false,
		CF: cf.CFConfig{
			SkipSSLValidation: false,
		},
		RateLimit: models.RateLimitConfig{
			MaxAmount:     DefaultMaxAmount,
			ValidDuration: DefaultValidDuration,
		},
	}

	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	err = yaml.UnmarshalStrict(bytes, conf)
	if err != nil {
		return nil, err
	}

	conf.Logging.Level = strings.ToLower(conf.Logging.Level)

	return conf, nil
}

func (c *Config) Validate() error {
	err := c.CF.Validate()
	if err != nil {
		return err
	}

	if c.DB.PolicyDB.URL == "" {
		return fmt.Errorf("Configuration error: PolicyDB URL is empty")
	}
	if c.Scheduler.SchedulerURL == "" {
		return fmt.Errorf("Configuration error: scheduler.scheduler_url is empty")
	}
	if c.ScalingEngine.ScalingEngineUrl == "" {
		return fmt.Errorf("Configuration error: scaling_engine.scaling_engine_url is empty")
	}
	if c.MetricsCollector.MetricsCollectorUrl == "" {
		return fmt.Errorf("Configuration error: metrics_collector.metrics_collector_url is empty")
	}
	if c.EventGenerator.EventGeneratorUrl == "" {
		return fmt.Errorf("Configuration error: event_generator.event_generator_url is empty")
	}
	if c.MetricsForwarder.MetricsForwarderUrl == "" {
		return fmt.Errorf("Configuration error: metrics_forwarder.metrics_forwarder_url is empty")
	}
	if c.PolicySchemaPath == "" {
		return fmt.Errorf("Configuration error: PolicySchemaPath is empty")
	}
	if c.RateLimit.MaxAmount <= 0 {
		return fmt.Errorf("Configuration error: RateLimit.MaxAmount is equal or less than zero")
	}
	if c.RateLimit.ValidDuration <= 0*time.Nanosecond {
		return fmt.Errorf("Configuration error: RateLimit.ValidDuration is equal or less than zero nanosecond")
	}
	if err := c.Health.Validate("api"); err != nil {
		return err
	}

	if c.InfoFilePath == "" {
		return fmt.Errorf("Configuration error: InfoFilePath is empty")
	}
	if !c.UseBuildInMode {
		if c.DB.BindingDB.URL == "" {
			return fmt.Errorf("Configuration error: BindingDB URL is empty")
		}

		for _, brokerCredential := range c.BrokerCredentials {
			if brokerCredential.BrokerUsername == "" && string(brokerCredential.BrokerUsernameHash) == "" {
				return fmt.Errorf("Configuration error: both broker_username and broker_username_hash are empty, please provide one of them")
			}
			if brokerCredential.BrokerUsername != "" && string(brokerCredential.BrokerUsernameHash) != "" {
				return fmt.Errorf("Configuration error: both broker_username and broker_username_hash are set, please provide only one of them")
			}
			if string(brokerCredential.BrokerUsernameHash) != "" {
				if _, err := bcrypt.Cost(brokerCredential.BrokerUsernameHash); err != nil {
					return fmt.Errorf("Configuration error: broker_username_hash is not a valid bcrypt hash")
				}
			}
			if brokerCredential.BrokerPassword == "" && string(brokerCredential.BrokerPasswordHash) == "" {
				return fmt.Errorf("Configuration error: both broker_password and broker_password_hash are empty, please provide one of them")
			}

			if brokerCredential.BrokerPassword != "" && string(brokerCredential.BrokerPasswordHash) != "" {
				return fmt.Errorf("Configuration error: both broker_password and broker_password_hash are set, please provide only one of them")
			}

			if string(brokerCredential.BrokerPasswordHash) != "" {
				if _, err := bcrypt.Cost(brokerCredential.BrokerPasswordHash); err != nil {
					return fmt.Errorf("Configuration error: broker_password_hash is not a valid bcrypt hash")
				}
			}
		}

		if c.CatalogSchemaPath == "" {
			return fmt.Errorf("Configuration error: CatalogSchemaPath is empty")
		}
		if c.CatalogPath == "" {
			return fmt.Errorf("Configuration error: CatalogPath is empty")
		}

		catalogSchemaLoader := gojsonschema.NewReferenceLoader("file://" + c.CatalogSchemaPath)
		catalogLoader := gojsonschema.NewReferenceLoader("file://" + c.CatalogPath)

		result, err := gojsonschema.Validate(catalogSchemaLoader, catalogLoader)
		if err != nil {
			return err
		}
		if !result.Valid() {
			errString := "{"
			for index, desc := range result.Errors() {
				if index == len(result.Errors())-1 {
					errString += fmt.Sprintf("\"%s\"", desc.Description())
				} else {
					errString += fmt.Sprintf("\"%s\",", desc.Description())
				}
			}
			errString += "}"
			return fmt.Errorf(errString)
		}
	}
	return nil
}
