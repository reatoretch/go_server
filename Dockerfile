FROM golang:latest

RUN mkdir /go/src/go_server

WORKDIR /go/src/go_server

RUN go get golang.org/x/lint/golint

ADD . /go/src/go_server

CMD ["go","run","/go/src/go_server/main.go"]
