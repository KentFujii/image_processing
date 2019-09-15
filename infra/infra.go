package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/s3/create_new_bucket_and_object.go
// https://aws.amazon.com/jp/blogs/developer/mocking-out-then-aws-sdk-for-go-for-unit-testing/
// https://dev.classmethod.jp/go/access-minio-using-aws-sdk-for-go/
// https://qiita.com/hmarf/items/7f4d39c48775c205b99b
func NewS3Infra(c s3Config) s3Infra {
	s, _ := session.NewSession()
	aak := c.ReadAwsAccountKey()
	ask := c.ReadAwsSecretKey()
	ar := c.ReadAwsRegion()
	ae := c.ReadAwsEndpoint()
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials(aak, ask, ""),
		Region: aws.String(ar),
		Endpoint: aws.String(ae),
		S3ForcePathStyle: aws.Bool(true),
	}
	return s3Infra{S3: s3.New(s, &cfg)}
}
