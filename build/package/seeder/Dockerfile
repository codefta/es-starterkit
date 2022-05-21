FROM golang:1.17.5-alpine3.15 as build
WORKDIR /go/src/github.com/ghazlabs/es-starterkit

COPY go.mod go.sum ./
RUN go mod download -x

COPY ./cmd/seeder ./cmd/seeder
COPY ./internal ./internal

WORKDIR /go/src/github.com/ghazlabs/es-starterkit/cmd/seeder
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o app

FROM alpine:3.15
RUN apk add ca-certificates tzdata
COPY --from=build /go/src/github.com/ghazlabs/es-starterkit/cmd/seeder .

ENTRYPOINT [ "./app" ]