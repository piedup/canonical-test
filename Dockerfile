FROM golang:1.9.2

RUN go get -u github.com/golang/dep/cmd/dep

COPY src $GOPATH/src/github.com/piedup/canonical-test

WORKDIR $GOPATH/src/github.com/piedup/canonical-test