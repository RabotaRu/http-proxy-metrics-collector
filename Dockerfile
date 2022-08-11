ARG GOLANG_VERSION=1.18
ARG ALPINE_VERSION=3.16

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV PROJECT=http-proxy-metrics-collector

WORKDIR ${PROJECT}

COPY go.mod go.sum ./
RUN go mod download

# Copy src code from the host and compile it
COPY . .
RUN go build -a -o /${PROJECT} .

### Base image with shell
FROM alpine:${ALPINE_VERSION} as base-release
RUN apk --update --no-cache add ca-certificates && update-ca-certificates
ENTRYPOINT ["/bin/http-proxy-metrics-collector"]

### Build with goreleaser
FROM base-release as goreleaser
COPY http-proxy-metrics-collector /bin/

### Build in docker
FROM base-release as release
COPY --from=builder /http-proxy-metrics-collector /bin/

### Scratch with build in docker
FROM scratch as scratch-release
COPY --from=base-release /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /http-proxy-metrics-collector /bin/
ENTRYPOINT ["/bin/http-proxy-metrics-collector"]
USER 65534

### Scratch with goreleaser
FROM scratch as scratch-goreleaser
COPY --from=base-release /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY http-proxy-metrics-collector /bin/
ENTRYPOINT ["/bin/http-proxy-metrics-collector"]
USER 65534
