package infra

import (
	// "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/common-examples.html
// https://godoc.org/github.com/aws/aws-sdk-go/service/s3/s3iface
// https://aws.amazon.com/jp/blogs/developer/mocking-out-then-aws-sdk-for-go-for-unit-testing/
type s3Infra struct {
	Client s3iface.S3API
	Bucket string
}

// func (i *s3Infra) CreateObject(key, content) {
// }

// func (i *s3Infra) ReadObject() {
// }
