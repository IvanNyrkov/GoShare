FROM golang

MAINTAINER Ivan Nyrkov ivan.nyrkov@gmail.com

WORKDIR /go/src/github.com/IvanNyrkov/GoShare
ADD . /go/src/github.com/IvanNyrkov/GoShare

ENV CONFIG_FILE=config-docker.json

RUN go build -o /go/src/github.com/IvanNyrkov/GoShare/build github.com/IvanNyrkov/GoShare/src
ENTRYPOINT /go/src/github.com/IvanNyrkov/GoShare/build

EXPOSE 80