FROM golang:1.15.7-alpine

ENV ROOT=/go/src/app
RUN apk update && apk add git
WORKDIR ${ROOT}

ADD . ${ROOT}

