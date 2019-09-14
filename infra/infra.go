package infra

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// type s3Config interface {
// 	ReadAwsAccountKey() string
// 	ReadAwsSecretKey() string
// 	ReadAwsRegion() string
// 	ReadAwsEndpoint() string
// 	ReadBucket() string
// }

// type awsS3 interface {
// }

// https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/s3/create_new_bucket_and_object.go
func NewS3Infra() s3Infra {
	s, _ := session.NewSession()
	ak := "image_processing"
	sk := "password"
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials(ak, sk, ""),
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("http://storage:9000"),
		S3ForcePathStyle: aws.Bool(true),
	}
	i := s3Infra{S3: s3.New(s, &cfg)}
	return i
}

type s3Infra struct {
	S3 *s3.S3
}

func (i *s3Infra) ReadObjects() {
	c := i.S3
	bucket := "image_processing"
	var token *string
	for complete := false; !complete; {
		in := s3.ListObjectsV2Input{Bucket: &bucket, ContinuationToken: token}
		out, err := c.ListObjectsV2(&in)
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

// https://dev.classmethod.jp/go/access-minio-using-aws-sdk-for-go/
// https://qiita.com/hmarf/items/7f4d39c48775c205b99b

