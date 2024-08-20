##############################################################
## Stage 1 - Go Build
##############################################################

FROM golang:1.21.6-bookworm AS builder

# Installing necessary packages
# make for Makefile support
# sqlite3 and libsqlite3-dev for SQLite support
# build-essential for common build requirements
RUN apt-get update && apt-get install -y --no-install-recommends \
    make gcc sqlite3 libsqlite3-dev build-essential && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/cloud-barista/cm-beetle

# Copy only necessary files
COPY go.mod go.sum go.work go.work.sum LICENSE Makefile ./
COPY api ./api
COPY cmd/cm-beetle ./cmd/cm-beetle
COPY conf ./conf
COPY pkg ./pkg
COPY scripts ./scripts

# NOTE - "make prod" executes the commannd, "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o cm-beetle"
RUN make prod

#############################################################
## Stage 2 - Application Setup
##############################################################

FROM ubuntu:22.04 AS prod

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

# Set the Current Working Directory inside the container
WORKDIR /app

## Copy the Pre-built binary and necessary files from the previous stage
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/scripts/ /app/scripts/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/conf/ /app/conf/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/cmd/cm-beetle /app/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/api/ /app/api/

## Set environment variables 
# Set system endpoints
ENV BEETLE_ROOT=/app \
    BEETLE_CBSTORE_ROOT=/app \
    BEETLE_CBLOG_ROOT=/app 

ENV BEETLE_TUMBLEBUG_REST_URL=http://localhost:1323/tumblebug

## Logger configuration
# Set log file path (default logfile path: ./beetle.log) 
# Set log level, such as trace, debug info, warn, error, fatal, and panic
ENV BEETLE_LOGFILE_PATH=/app/log/beetle.log \
    BEETLE_LOGFILE_MAXSIZE=10 \
    BEETLE_LOGFILE_MAXBACKUPS=3 \
    BEETLE_LOGFILE_MAXAGE=30 \
    BEETLE_LOGFILE_COMPRESS=false \
    BEETLE_LOGLEVEL=info

# Set execution environment, such as development or production
ENV BEETLE_NODE_ENV=production

## Set internal DB config (SQLlite)
ENV BEETLE_SQLITE_URL=localhost:3306 \
    BEETLE_SQLITE_DATABASE=cm_beetle \
    BEETLE_SQLITE_USER=cm_beetle \
    BEETLE_SQLITE_PASSWORD=cm_beetle

## Set API access config
# API_ALLOW_ORIGINS (ex: https://cloud-barista.org,xxx.xxx.xxx.xxx or * for all)
# Set ENABLE_AUTH=true currently for basic auth for all routes (i.e., url or path)
ENV BEETLE_API_ALLOW_ORIGINS=* \
    BEETLE_API_AUTH_ENABLED=true \
    BEETLE_API_USERNAME=default \
    BEETLE_API_PASSWORD=default

## Set period for auto control goroutine invocation
ENV BEETLE_AUTOCONTROL_DURATION_MS=10000

## Set SELF_ENDPOINT, to access Swagger API dashboard outside (Ex: export SELF_ENDPOINT=x.x.x.x:8056)
ENV BEETLE_SELF_ENDPOINT=localhost:8056

ENTRYPOINT [ "/app/cm-beetle" ]

EXPOSE 8056
