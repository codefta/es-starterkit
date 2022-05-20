FROM golang:1.17.5-alpine3.15 as build
WORKDIR /go/src/github.com/ghazlabs/es-starterkit

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./cmd/kuliner-api ./cmd/kuliner-api
COPY ./internal/kuliner-api ./internal/kuliner-api

WORKDIR /go/src/github.com/ghazlabs/es-starterkit/cmd/kuliner-api
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

FROM alpine:3.15
RUN apk add ca-certificates tzdata
COPY --from=build /go/src/github.com/ghazlabs/es-starterkit/cmd/kuliner-api .

ENTRYPOINT [ "./app" ]