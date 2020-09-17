GOCMD := go
GOBUILD := ${GOCMD} build
BIN_NAME := core-api

PROJECT_ROOT := ${shell pwd}
PROJECT_MAIN_PKG := cmd/core-api
PROJECT_DEBUG_OUTPUT := ${PROJECT_ROOT}/bin/debug
PROJECT_RELEASE_OUTPUT := ${PROJECT_ROOT}/bin/release
PROJECT_ENV_EXAMPLE := ${PROJECT_ROOT}/config/env-example
PROJECT_ENV_FILE := ${PROJECT_ROOT}/config/.env.yml

.PHONY: test
test:
	@-echo "  > Running test..."
	@$(GOCMD) test -v ./...

.PHONY: copy-env-debug
copy-env-debug:
	@-echo "  > Copying env-example to .env.yml...\n"
	@-cp ${PROJECT_ENV_EXAMPLE} ${PROJECT_ENV_FILE}
	@-echo " > Please edit ${PROJECT_ENV_FILE} with your server configuration."

.PHONY: compile
compile:
	@$(GOBUILD) -o $(PROJECT_DEBUG_OUTPUT)/$(BIN_NAME) $(PROJECT_ROOT)/$(PROJECT_MAIN_PKG)

.PHONY: run
run: compile
	@-echo "  > Starting Server...\n"
	@${PROJECT_DEBUG_OUTPUT}/${BIN_NAME}
