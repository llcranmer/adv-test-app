FROM golang:1.8.1-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git

RUN go get github.com/hashicorp/consul/api

ENV SOURCES /go/src/github.com/llcranmer/adv-test-app/Discovery/Simple/client

COPY . ${SOURCES}

RUN cd ${SOURCES}

ENV CONSUL_HTTP_ADDR localhost:8500

WORKDIR ${SOURCES}client/
CMD ${SOURCES}client/client
