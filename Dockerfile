##############################################################
## Stage 1 - Go Build
##############################################################

FROM golang:1.19-alpine AS builder

#RUN apk update && apk add --no-cache bash

#RUN apk add gcc

RUN apk add --no-cache sqlite-libs sqlite-dev

RUN apk add --no-cache build-base

ADD . /go/src/github.com/cloud-barista/cm-beetle

WORKDIR /go/src/github.com/cloud-barista/cm-beetle/pkg

RUN go build -ldflags '-w -extldflags "-static"' -tags cm-beetle -o cm-beetle -v

#############################################################
## Stage 2 - Application Setup
##############################################################

FROM ubuntu:latest as prod

# use bash
RUN rm /bin/sh && ln -s /bin/bash /bin/sh

WORKDIR /app/src

# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/assets/ /app/assets/

COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/scripts/ /app/scripts/

# COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/conf/ /app/conf/

COPY --from=builder /go/src/github.com/cloud-barista/cm-beetle/pkg/cm-beetle /app/pkg/

#RUN /bin/bash -c "source /app/conf/setup.env"
ENV CMBEETLE_ROOT /app
ENV CBSTORE_ROOT /app
ENV CBLOG_ROOT /app
# ENV SPIDER_CALL_METHOD REST
# ENV SPIDER_REST_URL http://cb-spider:1024/spider

ENV DB_URL localhost:3306
ENV DB_DATABASE cm_beetle
ENV DB_USER cm_beetle
ENV DB_PASSWORD cm_beetle

# API Setting
# ALLOW_ORIGINS (ex: https://cloud-barista.org,xxx.xxx.xxx.xxx or * for all)
ENV ALLOW_ORIGINS *
ENV SKIP_BASIC_AUTH false
ENV API_USERNAME default
ENV API_PASSWORD default

# Set period for auto control goroutine invocation
ENV AUTOCONTROL_DURATION_MS 10000

# Set SELF_ENDPOINT, if you want to access Swagger API dashboard from outside. (Ex: export SELF_ENDPOINT=xxx.xxx.xxx.xxx:8056)
ENV SELF_ENDPOINT localhost:8056

# Environment variables that you don't need to touch

# Swagger UI API document file path 
ENV API_DOC_PATH /app/pkg/api/rest/docs/swagger.json

ENTRYPOINT [ "/app/pkg/cm-beetle" ]

EXPOSE 8056
