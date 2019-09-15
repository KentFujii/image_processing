package infra

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/common-examples.html
// https://godoc.org/github.com/aws/aws-sdk-go/service/s3/s3iface
// https://aws.amazon.com/jp/blogs/developer/mocking-out-then-aws-sdk-for-go-for-unit-testing/
type s3Infra struct {
	Client s3iface.S3API
	Bucket string
}

func (i *s3Infra) ReadObjects() {
	var token *string
	for complete := false; !complete; {
		in := s3.ListObjectsV2Input{Bucket: &i.Bucket, ContinuationToken: token}
		out, err := i.Client.ListObjectsV2(&in)
		if err != nil {
			panic(err)
		}

		for i, o := range out.Contents {
			fmt.Printf("[%d] : %s\n", i, *o.Key)
		}

		complete = out.IsTruncated != nil && !*out.IsTruncated
		token = out.NextContinuationToken
	}
}
