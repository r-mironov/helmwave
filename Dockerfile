ARG GOLANG_VERSION=1.17
ARG ALPINE_VERSION=3.14
ARG PROJECT=helmwave

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder

LABEL maintainer="zhilyaev.dmitriy+${PROJECT}@gmail.com"
LABEL name=${PROJECT}

# enable Go modules support
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR ${PROJECT}

COPY go.mod go.sum ./
RUN go mod download

# Copy src code from the host and compile it
COPY cmd cmd
COPY pkg pkg
RUN go build -a -o /${PROJECT} ./cmd/${PROJECT}

###
FROM alpine:${ALPINE_VERSION} as base-release
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/bin/${PROJECT}"]

###
FROM base-release as goreleaser
COPY ${PROJECT} /bin/

###
FROM base-release
COPY --from=builder /${PROJECT} /bin/

###
FROM scratch
COPY ${PROJECT} /bin/
ENTRYPOINT ["/bin/${PROJECT}"]

