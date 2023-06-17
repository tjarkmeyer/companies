PROJECT_NAME := "companies"
PKG := "github.com/tjarkmeyer/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/...)
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)

.PHONY: dep build clean test coverage coverhtml lint

lint: ## Lint the files
	golangci-lint run -v

test: ## Run unittests
	@go test ./... -covermode=atomic -coverpkg=./... -coverprofile ./coverage.out

race: dep ## Run data race detector
	@go test -race -short ${PKG_LIST}

coverage: ## Generate global code coverage report
	./scripts/code-coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./scripts/code-coverage-html.sh;

dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u golang.org/x/lint/golint

build: dep ## Build the binary file
	@go build -tags 'cgo_off' -a -o $(PROJECT_NAME) ./cmd

run-dev: ## Run the project 
	go run ./cmd/main.go

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

mockgen: ## Generate a repository mock
	@mockgen -source=internal/v1/abstraction.go -destination=internal/v1/repositories/mocks/${PROJECT_NAME}_mock.go -package=mock_internal ${PKG} CompaniesRepository

start-dev: ## start development environment
	bash ./scripts/start-environment.sh

stop-dev: ## stop development environment
	bash ./scripts/stop-environment.sh

postgres-login: ## login to local postgres instance
	psql -h localhost -U postgres

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
