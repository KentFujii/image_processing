FROM golang:1.12.9

ENV GO111MODULE on
RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u github.com/oxequa/realize
RUN go get github.com/onsi/ginkgo/ginkgo
WORKDIR /go/src
