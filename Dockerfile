FROM golang:alpine as builder
WORKDIR /tui
COPY ["main.go", "go.mod", "go.sum", "./"]
COPY pkg/ pkg/
COPY vendor/ vendor/
RUN go build -mod=vendor

FROM alpine
COPY --from=builder /tui/tui /
ENTRYPOINT ["/tui"]
