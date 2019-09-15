package infra

import (
	"fmt"
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

var _ = Describe("S3", func() {
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
	Context("LoadS3Config", func() {
		It("Should create s3 object", func() {
			fmt.Println(s3Config)
			i := NewS3Infra(&s3Config)
			i.ReadObjects()
		})
	})
})
