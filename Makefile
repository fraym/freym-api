help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\Targets:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m	 %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: proto
proto: ## Generates api code from .proto files
	cd ./go && ./proto.sh
	cd ./js && npm install && npm run generate

lint: ## Run linters
	cd ./go && golangci-lint run --enable gofumpt
	cd ./js && npm run lint

test: ## Run tests
	cd ./go && go test ./...
