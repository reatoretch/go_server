FROM golang:latest

RUN mkdir /go/src/go_server

WORKDIR /go/src/go_server

ADD . /go/src/go_server
