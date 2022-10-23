.DEFAULT_GOAL := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

TARGET_ENTITY ?= TestEntity

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

.PHONY: ent-init
ent-init: ## Define a new entity of TARGET_ENTITY.
	go run -mod=mod entgo.io/ent/cmd/ent init ${TARGET_ENTITY}

.PHONY: ent-gen
ent-gen: ## Generate codes from entity models.
	go generate ./ent

.PHONY: ent-describe
ent-describe: ## Describe schemas of entity models.
	go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema

.PHONY: run
run: ## Run server locally.
	go run ./cmd
