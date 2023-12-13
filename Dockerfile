##############################################################
## Stage 1 - Go Build
##############################################################

FROM golang:1.21-alpine AS builder

#RUN apk update && apk add --no-cache bash

#RUN apk add gcc

RUN apk add --no-cache sqlite-libs sqlite-dev

RUN apk add --no-cache build-base

ADD . /go/src/github.com/cloud-barista/cm-beetle

WORKDIR /go/src/github.com/cloud-barista/cm-beetle/cmd/cm-beetle

RUN go build -ldflags '-w -extldflags "-static"' -tags cm-beetle -o cm-beetle -v

#############################################################
## Stage 2 - Application Setup
##############################################################

FROM ubuntu:latest as prod

# use bash
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

WORKDIR /app

# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/assets/ /app/assets/

COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/scripts/ /app/scripts/

# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/conf/ /app/conf/

COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/cmd/cm-beetle /app/

#RUN /bin/bash -c "source /app/conf/setup.env"
## Logger configuration
# Set log file path (default ./cm-beetle.log) 
ENV CM_BEETLE_LOG_PATH cm-beetle.log
ENV CM_BEETLE_LOG_MAX_SIZE 10
ENV CM_BEETLE_LOG_MAX_BACKUPS 3
ENV CM_BEETLE_LOG_MAX_AGE 30
ENV CM_BEETLE_LOG_COMPRESS false
# Set execution environment, such as development or production
ENV CM_BEETLE_APP_ENV production
# Set log level, such as trace, debug info, warn, error, fatal, and panic
ENV CM_BEETLE_LOG_LEVEL info

## Set system endpoints
ENV CMBEETLE_ROOT /app
ENV CBSTORE_ROOT /app
ENV CBLOG_ROOT /app
# ENV SPIDER_CALL_METHOD REST
# ENV SPIDER_REST_URL http://cb-spider:1024/spider

## Set internal DB config (SQLlite)
ENV DB_URL localhost:3306
ENV DB_DATABASE cm_beetle
ENV DB_USER cm_beetle
ENV DB_PASSWORD cm_beetle

## Set API access config
# ALLOW_ORIGINS (ex: https://cloud-barista.org,xxx.xxx.xxx.xxx or * for all)
ENV ALLOW_ORIGINS *
# Set ENABLE_AUTH=true currently for basic auth for all routes (i.e., url or path)
ENV ENABLE_AUTH true
ENV API_USERNAME default
ENV API_PASSWORD default

## Set period for auto control goroutine invocation
ENV AUTOCONTROL_DURATION_MS 10000

## Set SELF_ENDPOINT, to access Swagger API dashboard outside (Ex: export SELF_ENDPOINT=x.x.x.x:8056)
ENV SELF_ENDPOINT localhost:8056

## Environment variables that you don't need to touch
# Swagger UI API document file path 
ENV API_DOC_PATH /app/pkg/api/rest/docs/swagger.json

ENTRYPOINT [ "/app/cm-beetle" ]

EXPOSE 8056
