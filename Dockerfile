##############################################################
## Stage 1 - Go Build
##############################################################

FROM golang:1.23.0-bookworm AS builder

ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/cloud-barista/cm-beetle

# When you need to build cm-beetle with pkgs being developed locally,
# you can use the following guide.
# For example, this is useful when you are developing cm-beetle and cb-tumblebug at the same time.
# 1. [On go.mod] Add `replace github.com/cloud-barista/cb-tumblebug => ../cb-tumblebug`
# 2. [On shell] Copy cb-tumblebug by `rsync -av --exclude container-volume/ --exclude .git/ ../cb-tumblebug/ ./cb-tumblebug`
# 3. [On Dockerfile] Add `COPY cb-tumblebug /go/src/github.com/cloud-barista/cb-tumblebug`
# 4. [On shell] Run `make compose`
# 5. !Importance! After building/testing, make sure to restore(remove/comment) the changes by step 1-3.
# COPY cb-tumblebug /go/src/github.com/cloud-barista/cb-tumblebug

# Copy dependency files to the container
COPY go.mod go.sum go.work go.work.sum LICENSE ./
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download

# Copy some necessary files to the container
COPY api ./api
COPY cmd/cm-beetle ./cmd/cm-beetle
COPY pkg ./pkg
COPY scripts ./scripts
# COPY conf ./conf

# Build the Go app
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    cd cmd/cm-beetle && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -tags cm-beetle -v -o cm-beetle main.go

#############################################################
## Stage 2 - Application Setup
##############################################################

FROM ubuntu:22.04 AS prod

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

# Set the Current Working Directory inside the container
WORKDIR /app

# Installing necessary packages and cleaning up
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*

## Copy the Pre-built binary and necessary files from the previous stage
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/scripts/ /app/scripts/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/cmd/cm-beetle/cm-beetle /app/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/api/ /app/api/
# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/conf/ /app/conf/

## Set environment variables 
# Set system endpoints
ENV BEETLE_ROOT=/app 

## Set SELF_ENDPOINT, to access Swagger API dashboard outside (Ex: export SELF_ENDPOINT=x.x.x.x:8056)
ENV BEETLE_SELF_ENDPOINT=localhost:8056

## Set API access config
# API_ALLOW_ORIGINS (ex: https://cloud-barista.org,xxx.xxx.xxx.xxx or * for all)
# Set ENABLE_AUTH=true currently for basic auth for all routes (i.e., url or path)
ENV BEETLE_API_ALLOW_ORIGINS=* \
    BEETLE_API_AUTH_ENABLED=true \
    BEETLE_API_USERNAME=default \
    BEETLE_API_PASSWORD=default

## Set internal DB config (lkvstore: local key-value store, default file path: ./db/beetle.db)
ENV BEETLE_LKVSTORE_PATH=/app/db/beetle.db

## Logger configuration
# Set log file path (default logfile path: ./beetle.log) 
# Set log level, such as trace, debug info, warn, error, fatal, and panic
ENV BEETLE_LOGFILE_PATH=/app/log/beetle.log \
    BEETLE_LOGFILE_MAXSIZE=1000 \
    BEETLE_LOGFILE_MAXBACKUPS=3 \
    BEETLE_LOGFILE_MAXAGE=30 \
    BEETLE_LOGFILE_COMPRESS=false \
    BEETLE_LOGLEVEL=info \
    BEETLE_LOGWRITER=both

# Set execution environment, such as development or production
ENV BEETLE_NODE_ENV=production

## Set period for auto control goroutine invocation
ENV BEETLE_AUTOCONTROL_DURATION_MS=10000

## Set Tumblebug access config
ENV BEETLE_TUMBLEBUG_REST_URL=http://localhost:1323 \
    BEETLE_TUMBLEBUG_API_USERNAME=default \
    BEETLE_TUMBLEBUG_API_PASSWORD=default

ENTRYPOINT [ "/app/cm-beetle" ]

EXPOSE 8056
