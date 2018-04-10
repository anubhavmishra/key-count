BUILD_ID := $(shell git rev-parse --short HEAD 2>/dev/null || echo no-commit-id)
IMAGE := anubhavmishra/key-count

.DEFAULT_GOAL := help
help: ## List all targets & descriptions
	@cat Makefile* | grep -E '^[a-zA-Z_-]+:.*?## .*$$' | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## Clean the build
	rm -rf ./build

deps: ## Get project dependencies
	go get .

build: ## Builds the Go service
	go build -o key-count .

build-service: ## Build a docker container for the Go service
	mkdir -p ./build/linux/amd64
	GOOS=linux GOARCH=amd64 go build -v -o ./build/linux/amd64/key-count .
	docker build -t $(IMAGE):$(BUILD_ID) .
	docker tag $(IMAGE):$(BUILD_ID) $(IMAGE):latest

push: ## Docker push the service images tagged 'latest' & 'BUILD_ID'
	docker push $(IMAGE):$(BUILD_ID)
	docker push $(IMAGE):latest

deps-test: ## Test dependencies
	go get -t

test: ## Run tests
	go test -v .

run: ## Build and run the project locally
	mkdir -p ./build
	go build -o ./build/key-count && ./build/key-count
