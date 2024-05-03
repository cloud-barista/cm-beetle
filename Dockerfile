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

# NOTE - "make prod" executes the commannd, "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s -w' -o cm-beetle"
RUN make prod

#############################################################
## Stage 2 - Application Setup
##############################################################

FROM ubuntu:22.04 as prod

# use bash
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

WORKDIR /app

# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/assets/ /app/assets/

COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/scripts/ /app/scripts/

# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/conf/ /app/conf/

COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/cmd/cm-beetle /app/
COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/api/ /app/api/

## Set system endpoints
ENV CMBEETLE_ROOT /app
ENV CBSTORE_ROOT /app
ENV CBLOG_ROOT /app
# ENV CBSPIDER_CALLMETHOD REST
# ENV CBSPIDER_REST_URL http://cb-spider:1024/spider
# ENV CBTUMBLEBUG_CALLMETHOD REST
# ENV CBTUMBLEBUG_REST_URL http://localhost:1323/tumblebug

#RUN /bin/bash -c "source /app/conf/setup.env"
## Logger configuration
# Set log file path (default logfile path: ./cm-beetle.log) 
ENV LOGFILE_PATH cm-beetle.log
ENV LOGFILE_MAXSIZE 10
ENV LOGFILE_MAXBACKUPS 3
ENV LOGFILE_MAXAGE 30
ENV LOGFILE_COMPRESS false
# Set log level, such as trace, debug info, warn, error, fatal, and panic
ENV LOGLEVEL info
# Set execution environment, such as development or production
ENV NODE_ENV production

## Set internal DB config (SQLlite)
ENV DB_URL localhost:3306
ENV DB_DATABASE cm_beetle
ENV DB_USER cm_beetle
ENV DB_PASSWORD cm_beetle

## Set API access config
# API_ALLOW_ORIGINS (ex: https://cloud-barista.org,xxx.xxx.xxx.xxx or * for all)
ENV API_ALLOW_ORIGINS *
# Set ENABLE_AUTH=true currently for basic auth for all routes (i.e., url or path)
ENV API_AUTH_ENABLED true
ENV API_USERNAME default
ENV API_PASSWORD default

## Set period for auto control goroutine invocation
ENV AUTOCONTROL_DURATION_MS 10000

## Set SELF_ENDPOINT, to access Swagger API dashboard outside (Ex: export SELF_ENDPOINT=x.x.x.x:8056)
ENV SELF_ENDPOINT localhost:8056

## Environment variables that you don't need to touch
# Swagger UI API document file path 
ENV APIDOC_PATH /app/api/swagger.json

ENTRYPOINT [ "/app/cm-beetle" ]

EXPOSE 8056
