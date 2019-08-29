FROM golang:1.12.9

RUN go get -u github.com/go-delve/delve/cmd/dlv
RUN go get -u github.com/oxequa/realize
RUN go get -u github.com/onsi/ginkgo/ginkgo
RUN go get -u github.com/onsi/gomega
RUN go get -u github.com/aws/aws-sdk-go
RUN go get -u github.com/disintegration/gift
RUN go get -u github.com/disintegration/imaging
WORKDIR /go/src
