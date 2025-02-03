help: ## Show this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"; printf "\Targets:\n"} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m	 %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

.PHONY: proto
proto: ## Generates api code from .proto files
	cd ./go/proto && ./proto.sh
	cd ./js && npm install && npm run proto

lint: ## Lint the code
	cd ./go/streams && golangci-lint run --enable gofumpt
	cd js && npm run lint

release:
	## todo: use the release.md file to generate the release notes
	## push all needed git tags
	## publish release on npmjs
	## publish release on github

