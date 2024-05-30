##############################################################
## Stage 1 - Go Build
##############################################################

FROM golang:1.21.6-alpine AS builder

# Installing necessary packages
# make for Makefile support
# sqlite-libs and sqlite-dev for SQLite support
# build-base for common build requirements
RUN apk add --no-cache make gcc sqlite-libs sqlite-dev build-base

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/cloud-barista/cm-beetle

# Copy only necessary files
COPY go.mod go.sum go.work go.work.sum LICENSE Makefile ./
COPY api ./api
COPY cmd/cm-beetle ./cmd/cm-beetle
# COPY conf ./conf
COPY pkg ./pkg
COPY scripts ./scripts

# NOTE - "make prod" executes the commannd, "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o cm-beetle"
RUN make prod

#############################################################
## Stage 2 - Application Setup
##############################################################

FROM ubuntu:22.04 as prod

RUN rm /bin/sh && ln -s /bin/bash /bin/sh

# Set the Current Working Directory inside the container
WORKDIR /app

## Copy the Pre-built binary and necessary files from the previous stage
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/scripts/ /app/scripts/
# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/conf/ /app/conf/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/cmd/cm-beetle /app/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/api/ /app/api/

## Set environment variables 
# Set system endpoints
ENV CMBEETLE_ROOT=/app \
    CBSTORE_ROOT=/app \
    CBLOG_ROOT=/app

# ENV CBSPIDER_CALLMETHOD REST \
#     CBSPIDER_REST_URL http://cb-spider:1024/spider

ENV CBTUMBLEBUG_CALLMETHOD=REST \
    CBTUMBLEBUG_REST_URL=http://localhost:1323/tumblebug

## Logger configuration
# Set log file path (default logfile path: ./cm-beetle.log) 
# Set log level, such as trace, debug info, warn, error, fatal, and panic
ENV LOGFILE_PATH=$CMBEETLE_ROOT/log/cm-beetle.log \
    LOGFILE_MAXSIZE=10 \
    LOGFILE_MAXBACKUPS=3 \
    LOGFILE_MAXAGE=30 \
    LOGFILE_COMPRESS=false \
    LOGLEVEL=info

# Set execution environment, such as development or production
ENV NODE_ENV=production

## Set internal DB config (SQLlite)
ENV DB_URL=localhost:3306 \
    DB_DATABASE=cm_beetle \
    DB_USER=cm_beetle \
    DB_PASSWORD=cm_beetle

## Set API access config
# API_ALLOW_ORIGINS (ex: https://cloud-barista.org,xxx.xxx.xxx.xxx or * for all)
# Set ENABLE_AUTH=true currently for basic auth for all routes (i.e., url or path)
ENV API_ALLOW_ORIGINS=* \
    API_AUTH_ENABLED=true \
    API_USERNAME=default \
    API_PASSWORD=default

## Set period for auto control goroutine invocation
ENV AUTOCONTROL_DURATION_MS=10000

## Set SELF_ENDPOINT, to access Swagger API dashboard outside (Ex: export SELF_ENDPOINT=x.x.x.x:8056)
ENV SELF_ENDPOINT=localhost:8056

## Environment variables that you don't need to touch
# Swagger UI API document file path 
ENV APIDOC_PATH=/app/api/swagger.json

ENTRYPOINT [ "/app/cm-beetle" ]

EXPOSE 8056
