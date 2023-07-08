ENTITY ?= User

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

graph: goda graphviz ## Draw dependency graph.
	$(GODA) graph -cluster -short github.com/isutare412/goarch/ioc/... | dot -Tsvg -o graph.svg

.PHONY: test
test: ginkgo ## Run tests.
	$(GINKGO) ./...

.PHONY: run
run: ## Run server.
	go run ./cmd/...

##@ Build Dependencies

## Location to install dependencies to
LOCALBIN := $(shell pwd)/bin
$(LOCALBIN):
	@mkdir -p $(LOCALBIN)

## Tool Binaries
GINKGO ?= $(LOCALBIN)/ginkgo
GODA ?= $(LOCALBIN)/goda

## Tool Versions
GINKGO_VERSION ?= v2.11.0

.PHONY: ginkgo
ginkgo: $(GINKGO) ## Install ginkgo locally if necessary.
$(GINKGO): $(LOCALBIN)
	@test -s $(LOCALBIN)/ginkgo || GOBIN=$(LOCALBIN) go install github.com/onsi/ginkgo/v2/ginkgo@$(GINKGO_VERSION)

.PHONY: goda
goda: $(GODA) ## Install goda locally if necessary.
$(GODA): $(LOCALBIN)
	@test -s $(LOCALBIN)/goda || GOBIN=$(LOCALBIN) go install github.com/loov/goda@latest

.PHONY: graphviz
graphviz: ## Install graphviz if necessary.
	@command -v dot > /dev/null || brew install graphviz
