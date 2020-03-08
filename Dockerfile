FROM golang:alpine as builder
WORKDIR /build
ADD . /build
RUN go version && \
    go env && \
    CGO_ENABLED=0 GOOS=linux go build

FROM alpine:latest
MAINTAINER Seth Hoenig <seth.a.hoenig@gmail.com>

WORKDIR /opt
COPY --from=builder /build/consul-socket /opt

ENTRYPOINT ["/opt/consul-socket"]

## Example Build
#     docker build -t shoenig/consul-socket .

## Example Publish
#     docker tag shoenig/consul-socket:latest shoenig/consul-socket:v0.0.0
#     docker push shoenig/consul-socket:v0.0.0

## Example launch (no task namespace)
#     docker run --rm -p 127.0.0.1:8500:8500/tcp -v /tmp/consul-test.sock:/tmp/consul-test.sock shoenig/consul-socket --socket=/tmp/consul-test.sock --bind 0.0.0.0:8500
