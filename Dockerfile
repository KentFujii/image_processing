FROM golang:1.12.9

RUN apt-get update && \
  apt-get install -y imagemagick libmagickcore-dev libmagickwand-dev && \
  apt-get clean

RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/onsi/ginkgo/ginkgo

ENV GO111MODULE on
WORKDIR /go/src
