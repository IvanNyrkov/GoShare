FROM golang

MAINTAINER Ivan Nyrkov ivan.nyrkov@gmail.com

WORKDIR /go/src/github.com/nrkv/GoShare
ADD . /go/src/github.com/nrkv/GoShare

ENV CONFIG_FILE=config-docker.json

RUN go build -o /go/src/github.com/nrkv/GoShare/build github.com/nrkv/GoShare
ENTRYPOINT /go/src/github.com/nrkv/GoShare/build

EXPOSE 80