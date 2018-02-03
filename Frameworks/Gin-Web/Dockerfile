FROM golang:1.9.2-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git 
RUN go get github.com/gin-gonic/gin

ENV SOURCES /go/src/github.com/llcranmer/adv-test-app/Frameworks/Gin-Web/

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}

CMD ${SOURCES}Gin-Web
EXPOSE 8080
