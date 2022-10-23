.DEFAULT_GOAL := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

PROJECT_NAME := goarch
MODE ?= local
COMPOSE_FILE := ${ROOTDIR}/deployments/docker-compose.infra.yaml
ENV_FILE := ${ROOTDIR}/deployments/${MODE}.env

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Infrastructure

.PHONY: up
up: ## Up infrastructures for development.
	docker-compose -f ${COMPOSE_FILE} --env-file ${ENV_FILE} -p ${PROJECT_NAME} up -d

.PHONY: down
down: ## Down infrastructures for development.
	docker-compose -f ${COMPOSE_FILE} --env-file ${ENV_FILE} -p ${PROJECT_NAME} down

.PHONY: ps
ps: ## Check processes of infrastructures for development.
	docker-compose -f ${COMPOSE_FILE} --env-file ${ENV_FILE} -p ${PROJECT_NAME} ps

.PHONY: logs
logs: ## Follow logs of infrastructures for development.
	docker-compose -f ${COMPOSE_FILE} --env-file ${ENV_FILE} -p ${PROJECT_NAME} logs -f
