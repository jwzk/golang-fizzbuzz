DOCKER_APP = fizzbuzz
DOCKER_BIN = docker-compose -p fizzbuzz
DOCKER_CMD = $(DOCKER_BIN) $(DOCKER_CONF)
DOCKER_CONF = -f docker-compose.yml
DOCKER_EXEC = $(DOCKER_CMD) exec $(DOCKER_APP)
DOCKER_UP_DETACH = $(DOCKER_CMD) up -d

##
## Tools
##

logs: ## App Logs
logs:
	@$(DOCKER_CMD) logs -f $(DOCKER_APP)

test: ## Run test
test:
	@$(DOCKER_EXEC) go test ./... -cover

test-cover: ## Run test with full coverage
test-cover:
	@$(DOCKER_EXEC) sh -c "go test ./... -coverprofile=/tmp/coverage.out && go tool cover -func=/tmp/coverage.out"

cli: ## App CLI with bash
cli:
	@$(DOCKER_EXEC) bash

.PHONY: logs test test-cover cli

##
## Installation
##

build: ## Build images
build:
	@$(DOCKER_CMD) build

up: ## Up containers
up:
	@$(DOCKER_UP_DETACH)

install: ## Build and Up containers
install: build up

update: ## Update App container
update:
	@$(DOCKER_CMD) build $(DOCKER_APP)
	@$(DOCKER_UP_DETACH) $(DOCKER_APP)

delete: ## Delete all containers
	@$(DOCKER_CMD) stop
	@$(DOCKER_CMD) rm --force

clean: ## Delete and install
clean: delete install

.PHONY: build up install update delete clean

##

.DEFAULT_GOAL := help
help:
	@grep -E '(^[a-zA-Z_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'
.PHONY: help
