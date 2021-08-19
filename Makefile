SHELL := /bin/bash

MVN_OPTS="-Dmaven.test.skip=true"
PROFILE=default
POSTGRES_URL=jdbc:postgresql://127.0.0.1/autoscaler
POSTGRES_CHANGELOGS=\
     src/autoscaler/api/db/api.db.changelog.yml \
     src/autoscaler/servicebroker/db/servicebroker.db.changelog.json \
     scheduler/db/scheduler.changelog-master.yaml  \
     scheduler/db/quartz.changelog-master.yaml \
     src/autoscaler/metricsserver/db/metricscollector.db.changelog.yml  \
     src/autoscaler/eventgenerator/db/dataaggregator.db.changelog.yml   \
     src/autoscaler/scalingengine/db/scalingengine.db.changelog.yml  \
     src/autoscaler/operator/db/operator.db.changelog.yml

init-postgres: build-db
	@for CHANGELOG in ${POSTGRES_CHANGELOGS}; \
		do java -cp 'db/target/lib/*' liquibase.integration.commandline.Main \
		--url=${POSTGRES_URL} --driver=org.postgresql.Driver \
		--username="postgres" --password="postgres"\
		--changeLogFile=$${CHANGELOG} update; done

githooks:
	@which pre-commit > /dev/null || brew install pre-commit
	@precommit install

init:
	@make -C src/autoscaler buildtools

build-db:
	mvn package -pl db ${MVN_OPTS}

build-scheduler:
	@mvn package -pl scheduler ${MVN_OPTS}

build-autoscaler:
	@make -C src/autoscaler build

.PHONY: test-certs
test-certs:
	@./scripts/generate_test_certs.sh
	@./scheduler/scripts/generate_unit_test_certs.sh

.PHONY: test
test:
	@make -C src/autoscaler test
	@mvn --no-transfer-progress test -Dspring.profiles.active=$PROFILE

.PHONY: start-postgres
start-postgres:
	@docker run -p 5432:5432  --name postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=autoscaler -d postgres:9.6

.PHONY: stop-postgres
stop-postgres:
	@docker rm -f postgres

build: init test-certs build-db build-scheduler build-autoscaler