.DEFAULT_GOAL := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

TARGET_ENTITY ?= TestEntity

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: ent-init
ent-init: ## Define a new entity of TARGET_ENTITY.
	go run -mod=readonly entgo.io/ent/cmd/ent --target pkg/core/ent/schema init ${TARGET_ENTITY}

.PHONY: ent-gen
ent-gen: ## Generate codes from entity models.
	go generate ./pkg/core/ent

.PHONY: ent-describe
ent-describe: ## Describe schemas of entity models.
	go run -mod=readonly entgo.io/ent/cmd/ent describe ./pkg/core/ent/schema

.PHONY: docs
docs: ## Generate OpenAPI specification for Swagger UI.
	go run -mod=readonly github.com/swaggo/swag/cmd/swag init \
		-d "${ROOTDIR}/cmd,${ROOTDIR}/pkg/controller/http,${ROOTDIR}/pkg/core/dto" \
		-o ${ROOTDIR}/api

.PHONY: run
run: ## Run server locally.
	go run ./cmd

.PHONY: test
test: ## Run tests.
	go test ./...
