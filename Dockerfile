FROM golang

MAINTAINER Ivan Nyrkov ivan.nyrkov@gmail.com

WORKDIR /go/src/github.com/IvanNyrkov/GoShare
ADD ./public /go/src/github.com/IvanNyrkov/GoShare/public
ADD ./src /go/src/github.com/IvanNyrkov/GoShare/src

# DEFAULT ENVIRONMENTS
ENV ENVIRONMENT=development
ENV EXPOSE_PORT=:80

# RUN PROJECT
RUN go build -o /go/src/github.com/IvanNyrkov/GoShare/build github.com/IvanNyrkov/GoShare/src
ENTRYPOINT /go/src/github.com/IvanNyrkov/GoShare/build

EXPOSE 80