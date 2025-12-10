ifeq ($(OS),Windows_NT)
    PWD := $(shell cd)
else
    PWD := $(shell pwd -L)
endif

ARCH := $(shell uname -m)
PLATFORM :=

ifeq ($(ARCH),arm64)
    PLATFORM := --platform=linux/amd64
endif

IMAGE = golang:1.25-alpine

GO_CACHE     := $(PWD)/.cache/go/build
GO_MOD_CACHE := $(PWD)/.cache/go/pkg/mod

DOCKER_RUN = docker run ${PLATFORM} --rm -it \
	-v ${PWD}:/app \
	-v ${GO_CACHE}:/root/.cache/go-build \
	-v ${GO_MOD_CACHE}:/go/pkg/mod \
	-w /app ${IMAGE} sh -c

.DEFAULT_GOAL := help

.PHONY: configure
configure: clean ## Configure development environment
	@mkdir -p .cache/go/build .cache/go/pkg/mod
	@${DOCKER_RUN} "go mod init capital-gains \
		&& go get -u -t ./... \
    	&& go mod tidy \
    	&& go mod vendor"

.PHONY: test
test: ## Run tests with coverage
	@mkdir -p reports/coverage \
		&& ${DOCKER_RUN} "go test -p=12 -parallel=6 \
			-coverprofile=reports/coverage/coverage.out -covermode=atomic ./src/application/... \
			&& go tool cover -html=reports/coverage/coverage.out -o reports/coverage/coverage.html"

.PHONY: review
review: ## Run static code analysis
	@${DOCKER_RUN}

.PHONY: show-reports
show-reports: ## Open static analysis reports (e.g., coverage, lints) in the browser
	@sensible-browser reports/coverage/coverage.html

.PHONY: run
run: ## Run the application
	@${DOCKER_RUN} "go run src/main.go"

.PHONY: clean
clean: ## Remove dependencies and generated artifacts
	@sudo chown -R ${USER}:${USER} ${PWD}
	@rm -rf go.mod go.sum vendor reports

.PHONY: help
help: ## Display this help message
	@echo "Usage: make [target]"
	@echo ""
	@echo "Setup and run"
	@grep -E '^(configure|run):.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Testing"
	@grep -E '^(test):.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Code review"
	@grep -E '^(review):.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Reports"
	@grep -E '^(show-reports):.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Cleanup"
	@grep -E '^(clean):.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "Help"
	@grep -E '^(help):.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'
