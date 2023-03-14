.PHONY: base build run clean

GIT_VERSION = $(shell describe --tags 2>/dev/null)
BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS = -X github.com/anaxaim/tui/pkg/version.gitVersion=$(GIT_VERSION) \
	-X github.com/anaxaim/tui/pkg/version.buildDate=$(BUILD_DATE)

GOFILES = $(shell find . -name "*.go" -type f -not -path "./vendor/*")

base: clean lint fmt build

clean: ## clean bin and go mod
	@rm -rf bin/
	go mod tidy
	go mod vendor

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

lint: install-golangci-lint ## run golangci lint
	golangci-lint run

fmt: ## golang format
	gofmt -s -w $(GOFILES)

build: ## build server
	go build -ldflags "$(LDFLAGS)" -mod vendor -o bin/tui main.go

run: ## run server
	go run -mod vendor main.go

SERVER_IMG=tui-server
docker-build-server: ## build server image
	docker build -t $(SERVER_IMG) .

docker-run-server: ## run server in docker
	docker run --network host -v $(shell pwd)/config:/config -v /var/run/docker.sock:/var/run/docker.sock $(SERVER_IMG)

mongo: ## init mongo db
	@docker start mymongo || docker run --name mymongo -d -p 27017:27017 mongo:6.0.4

