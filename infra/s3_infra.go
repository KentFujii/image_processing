package infra

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type s3Config interface {
	ReadAwsAccountKey() string
	ReadAwsSecretKey() string
	ReadAwsRegion() string
	ReadAwsEndpoint() string
	ReadBucket() string
}

type s3Infra struct {
	// S3 *s3.S3
	S3 s3iface.S3API
}

func (i *s3Infra) ReadObjects() {
	bucket := "image_processing"
	var token *string
	for complete := false; !complete; {
		in := s3.ListObjectsV2Input{Bucket: &bucket, ContinuationToken: token}
		out, err := i.S3.ListObjectsV2(&in)
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
