help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\Targets:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m	 %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

lint: ## Run linters
	golangci-lint run

format: ## Run formatters
	golangci-lint fmt

test: ## Run tests
	go test -short ./...

cover: ## Calculate test coverage
	go test -short -coverprofile=coverage.out ./... -timeout 2m
	go tool cover -html=coverage.out
