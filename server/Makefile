GIT_VERSION = $(shell describe --tags 2>/dev/null)
BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS = -X github.com/anaxaim/tui/server/pkg/version.gitVersion=$(GIT_VERSION) \
	-X github.com/anaxaim/tui/server/pkg/version.buildDate=$(BUILD_DATE)

GOFILES = $(shell find . -name "*.go" -type f -not -path "./vendor/*")

clean: ## clean bin and go mod
	@rm -rf ../bin/
	go mod tidy
	go mod vendor

fmt: ## golang format
	gofmt -s -w $(GOFILES)

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.51.2

lint: fmt install-golangci-lint ## run golangci lint
	golangci-lint run

run: ## run server
	go run -mod vendor main.go

build: ## build server
	go build -ldflags "$(LDFLAGS)" -mod vendor -o ../bin/tui main.go

SERVER_IMG=tui-server
MONGO_IMG=tui-mongo
mongo:
	docker rm -f $(MONGO_IMG) &>/dev/null && echo 'Removed old container' ; \
	docker run --name $(MONGO_IMG) -d -p 27017:27017 mongo:6.0.5

docker-build-server: ## build server image
	docker build -t $(SERVER_IMG) .

docker-run-server:
	docker rm -f $(SERVER_IMG) &>/dev/null && echo 'Removed old container' ; \
	docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8888:8888 -e MONGO_HOST=$(MONGO_IMG) --name $(SERVER_IMG) $(SERVER_IMG)
