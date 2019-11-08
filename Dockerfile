FROM golang:latest
RUN mkdir -p /go/src/github.com/pashukhin/indriver-test-task
ADD . /go/src/github.com/pashukhin/indriver-test-task
WORKDIR /go/src/github.com/pashukhin/indriver-test-task
RUN go get -v
