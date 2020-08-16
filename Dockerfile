FROM golang:latest

ARG UID=1001
RUN useradd -m -u ${UID} docker

USER ${UID}

RUN mkdir /go/src/go_server

WORKDIR /go/src/go_server

RUN go get golang.org/x/lint/golint

ADD . /go/src/go_server

CMD ["go","run","/go/src/go_server/main.go"]
