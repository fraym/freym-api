help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\Targets:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m	 %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

dev: ## Runs the dev environment on the k8s cluster
	kubectl apply -f .k8s/1-service.yml
	kubectl apply -f .k8s/2-deployment.yml
	okteto up -n fraym -f .okteto.yml

dev-stop: ## Removes the dev environment from the k8s cluster
	okteto down -n fraym -f .okteto.yml
	kubectl delete -f .k8s/2-deployment.yml
	kubectl delete -f .k8s/1-service.yml

lint: ## Run linters
	golangci-lint run --enable gofumpt

test: ## Run tests
	go test ./...

cover: ## Calculate test coverage
	go test -coverprofile=coverage.out ./... -timeout 2m
	go tool cover -html=coverage.out
