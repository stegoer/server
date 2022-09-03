FROM golang:1.17 as builder

WORKDIR /build

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s" -v -o stegoer cmd/server/server.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /build/stegoer ./stegoer

ENTRYPOINT ["./stegoer"]
