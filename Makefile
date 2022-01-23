.PHONY: all

all: slvlib

dep: ## Get the dependencies
	@go get -v -d ./...

slvlib: ## Build slvlib
	@CGO_ENABLED=0 GOOS=linux go build -v ./...

help: ## help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
