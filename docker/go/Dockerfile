FROM golang:alpine

ENV GO111MODULE=on

WORKDIR /go/src/app

COPY ./app .

ENV GOPATH /go/src

RUN apk update && \
    apk add vim curl git && \
    go get github.com/gin-contrib/cors && \
    go get github.com/gin-gonic/gin && \
    go get github.com/go-sql-driver/mysql&& \
    go get github.com/gorilla/handlers && \
    go get github.com/jinzhu/gorm

EXPOSE 8080
CMD ["go","run","main.go"]

