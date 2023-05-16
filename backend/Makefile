GIT_VERSION = $(shell describe --tags 2>/dev/null)
BUILD_DATE = $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')

LDFLAGS = -X github.com/anaxaim/tui/backend/pkg/version.gitVersion=$(GIT_VERSION) \
	-X github.com/anaxaim/tui/backend/pkg/version.buildDate=$(BUILD_DATE)

GOFILES = $(shell find . -name "*.go" -type f -not -path "./vendor/*")

clean: ## clean bin and go mod
	@rm -rf ../bin/
	go mod tidy
	go mod vendor

# CHECKS
vet:
	go vet ./...

fmt:
	go fmt ./...

gofumpt:
	go install mvdan.cc/gofumpt@v0.4.0
	gofumpt -l -w .

checks: vet fmt gofumpt

install-golangci-lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

linters: install-golangci-lint ## run golangci lint
	golangci-lint run

run: ## run server
	go run -mod vendor cmd/server/main.go

build: ## build server
	go build -ldflags "$(LDFLAGS)" -mod vendor -o ../bin/tui cmd/server/main.go


MONGO_IMG=tui-mongo
mongo:
	docker rm -f $(MONGO_IMG) &>/dev/null && echo 'Removed old mongo container' ; \
	docker run --name $(MONGO_IMG) -d -p 27017:27017 mongo:6.0.5

REDIS_IMG=tui-redis
redis:
	docker rm -f $(REDIS_IMG) &>/dev/null && echo 'Removed old redis container' ; \
    docker run --name $(REDIS_IMG) -d -p 6379:6379 redis:7-alpine

SERVER_IMG=tui-server
docker-build-server:
	docker build -t $(SERVER_IMG) .

docker-run-server:
	docker rm -f $(SERVER_IMG) &>/dev/null && echo 'Removed old server container' ; \
	docker run -d -v /var/run/docker.sock:/var/run/docker.sock -p 8888:8888 -e MONGO_HOST=$(MONGO_IMG) --name $(SERVER_IMG) $(SERVER_IMG)
