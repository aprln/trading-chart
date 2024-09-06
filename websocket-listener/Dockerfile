ARG GOLANG_VERSION=1.22
ARG ALPINE_VERSION=3.20

# Step 1: Build binaries
FROM golang:1.22-alpine${ALPINE_VERSION} AS builder
RUN apk add build-base

RUN mkdir -p /src
WORKDIR  /src
ADD ../.. .

# wslistener
RUN go mod tidy
RUN go build -o ./bin/wslistener ./websocket-listener

# Step 2: Build docker image
FROM alpine:${ALPINE_VERSION}
WORKDIR  /

COPY --from=builder /src/bin/wslistener /wslistener

CMD ["/wslistener"]