ARG GOLANG_VERSION=1.16
ARG ALPINE_VERSION=3.14
ARG DOCKER_VERSION=20.10.7

FROM golang:${GOLANG_VERSION}-alpine${ALPINE_VERSION} AS builder

LABEL maintainer="zhilyaev.dmitriy+biamon@gmail.com"
LABEL name="biamon"

# enable Go modules support
ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR biamon

COPY go.mod go.sum ./
RUN go mod download

# Copy src code from the host and compile it
COPY cmd cmd
COPY pkg pkg
RUN go build -a -o /biamon .

###
FROM docker:${DOCKER_VERSION}-git as base-release
RUN apk --no-cache add ca-certificates
ENTRYPOINT ["/bin/biamon"]

###
FROM base-release as goreleaser
COPY biamon /bin/

###
FROM base-release
COPY --from=builder /biamon /bin/