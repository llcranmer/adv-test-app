FROM golang:1.9.2-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

ENV SOURCES /go/src/github.com/llcranmer/adv-test-app/Discovery/Kubernetes/client/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}Client
