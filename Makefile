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

dev-go: ## Runs the go dev environment on the k8s cluster
	kubectl apply -f .k8s/1-go-service.yml
	kubectl apply -f .k8s/2-go-deployment.yml
	okteto up -n fraym -f .okteto.go.yml

dev-go-stop: ## Removes the go dev environment from the k8s cluster
	okteto down -n fraym -f .okteto.go.yml
	kubectl delete -f .k8s/2-go-deployment.yml
	kubectl delete -f .k8s/1-go-service.yml

dev-js: ## Runs the js dev environment on the k8s cluster
	kubectl apply -f .k8s/1-js-service.yml
	kubectl apply -f .k8s/2-js-deployment.yml
	okteto up -n fraym -f .okteto.js.yml

dev-js-stop: ## Removes the js dev environment from the k8s cluster
	okteto down -n fraym -f .okteto.js.yml
	kubectl delete -f .k8s/2-js-deployment.yml
	kubectl delete -f .k8s/1-js-service.yml
