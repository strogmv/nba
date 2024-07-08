.PHONY: help

#########
## DEV ##
#########

pre-commit-check: ## run all pre commit checks
pre-commit-check: go-mod-tidy go-generate go-fmt go-test go-test

commit-check: ## run pre commit checks and integration tests
commit-check: pre-commit-check run-it-tests

########
## GO ##
########
export GOPATH=$(shell go env GOPATH)

go-generate:	## run go generate command
	@go generate ./...

go-fmt:	## run go fmt over the project
	@go fmt ./...

go-lint:	## run go linters over the project
	@golangci-lint run ./...

go-test:	## run unit tests
	@go test $$(go list ./... | grep -v /test/)

go-download: 	## download all go dependencies
	@go mod download -x

go-mod-tidy: go-download 	## synchronize the go workspace (go.work) and go.mod/go.sum in each go package in the repository
	@go mod tidy

########################
## DOCKER ENVIRONMENT ##
########################

run-env:	## run environment
	@docker compose up dependent

run-server:	## run environment
	@docker compose up server

stop-env:	## stop environment
	@docker compose down \
		--volumes \
		--remove-orphans

run-all:	## run all
	@docker compose up

#######################
## INTEGRATION TESTS ##
#######################

run-it-tests: stop-env ## run environment
	@docker compose up tester

##########
## MISC ##
##########

install-go-tools:	## install go tools
	@go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v2.1.0
	@go install github.com/onsi/ginkgo/v2/ginkgo@v2.17.3
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.58.1



help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-43s\033[0m %s\n", $$1, $$2}'