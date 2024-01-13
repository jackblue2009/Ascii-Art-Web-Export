# syntax=docker/dockerfile:1
FROM golang:1.20-alpine

WORKDIR /app

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin main.go

LABEL unwanted=/tmp
LABEL Institute="Reboot01"
LABEL version="1.0"
LABEL description="Ascii-Art-Web Export Project"

EXPOSE 8080:8080

CMD ["go", "run", "main.go"]
