FROM golang:alpine AS server

WORKDIR /app
COPY . /app/

RUN cd /app && go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/tui main.go

FROM alpine:3.17

RUN apk add --no-cache \
    curl \
    docker-cli \
    docker-compose

WORKDIR /app

COPY --from=server /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=server /app/bin/tui /app/
COPY --from=server /app/internal/migrations/*.json /app/internal/migrations/
COPY --from=server /app/config/app.yaml /app/config/

ENTRYPOINT [ "/app/tui" ]
