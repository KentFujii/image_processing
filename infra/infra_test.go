package infra

import (
	"fmt"
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"bytes"
	// "strings"
)

type mockS3Config struct {
	AwsAccountKey string
	AwsSecretKey string
	AwsRegion string
	AwsEndpoint string
	Bucket string
}

func (c *mockS3Config) ReadAwsAccountKey() string {
	return c.AwsAccountKey
}

func (c *mockS3Config) ReadAwsSecretKey() string {
	return c.AwsSecretKey
}

func (c *mockS3Config) ReadAwsRegion() string {
	return c.AwsRegion
}

func (c *mockS3Config) ReadAwsEndpoint() string {
	return c.AwsEndpoint
}

func (c *mockS3Config) ReadBucket() string {
	return c.Bucket
}

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infra Suite")
}

var _ = Describe("Infra", func() {
	var s3Config mockS3Config
	BeforeEach(func() {
		s3Config = mockS3Config{
			AwsAccountKey: "image_processing",
			AwsSecretKey: "password",
			AwsRegion: "ap-northeast-1",
			AwsEndpoint: "http://storage:9000",
			Bucket: "image_processing",
		}

	})
	Context("NewS3Infra", func() {
		// https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/s3
		It("Should crud s3 object", func() {
			// create
			i := NewS3Infra(&s3Config)
			params := &s3.PutObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test.txt"),
				Body: bytes.NewReader([]byte("test!!!")),
				ContentType: aws.String("text/plain"),
			}
			resp, _ := i.Client.PutObject(params)
			fmt.Println(resp)
		})
	})
})
