ARG GOLANG_VERSION=1.22
ARG ALPINE_VERSION=3.20

# Step 1: Build binaries
FROM golang:1.22-alpine${ALPINE_VERSION} AS builder
RUN apk add build-base

RUN mkdir -p /src
WORKDIR  /src
ADD ../.. .

# grpcserver
RUN go mod tidy
RUN go build -o ./bin/grpcserver ./grpc-server

# migrate
RUN GOBIN=/src/bin go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Step 2: Build docker image
FROM alpine:${ALPINE_VERSION}
WORKDIR  /

COPY --from=builder /src/bin/grpcserver /grpcserver
COPY --from=builder /src/grpc-server/db /db
COPY --from=builder /src/bin/migrate /migrate

CMD ["./grpcserver"]