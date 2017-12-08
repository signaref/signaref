FROM golang

ADD . /go/src/

RUN go install UM-Server

ENTRYPOINT /go/bin/UM-Server

EXPOSE 9392
