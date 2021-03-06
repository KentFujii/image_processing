package infra

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"bytes"
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

type mockImageMagickConfig struct {
	ConvertTo string
	FormatWhitelist []string
	ResizeToLimit map[string]int
	ResizeToFit map[string]int
}

func (c *mockImageMagickConfig) ReadConvertTo() string {
	return c.ConvertTo
}
func (c *mockImageMagickConfig) ReadFormatWhitelist() []string {
	return c.FormatWhitelist
}

func (c *mockImageMagickConfig) ReadResizeToLimit() map[string]int {
	return c.ResizeToLimit
}

func (c *mockImageMagickConfig) ReadResizeToFit() map[string]int {
	return c.ResizeToFit
}

func TestInfra(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infra Suite")
}

var _ = Describe("Infra", func() {
	var s3Config mockS3Config
	var imageMagickConfig mockImageMagickConfig
	BeforeEach(func() {
		s3Config = mockS3Config{
			AwsAccountKey: "image_processing",
			AwsSecretKey: "password",
			AwsRegion: "ap-northeast-1",
			AwsEndpoint: "http://storage:9000",
			Bucket: "image_processing",
		}
		imageMagickConfig = mockImageMagickConfig{
			ConvertTo: "jpeg",
			FormatWhitelist: []string{"jpeg", "gif", "png"},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
			ResizeToFit: map[string]int{"height": 100, "width": 100},
		}
	})
	Context("NewS3Infra", func() {
		It("Should Put/List/Delete s3 object", func() {
			i := NewS3Infra(&s3Config)
			var putObjectParams *s3.PutObjectInput
			var err error
			// Put
			putObjectParams = &s3.PutObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test/test.txt"),
				Body: bytes.NewReader([]byte("test!")),
				ContentType: aws.String("text/plain"),
			}
			_, err = i.Client.PutObject(putObjectParams)
			Expect(err).To(BeNil())
			putObjectParams = &s3.PutObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test/test.txt"),
				Body: bytes.NewReader([]byte("test!!")),
				ContentType: aws.String("text/plain"),
			}
			_, err = i.Client.PutObject(putObjectParams)
			Expect(err).To(BeNil())
			// List
			listObjectsParams := &s3.ListObjectsInput{
				Bucket: aws.String(i.Bucket),
				Prefix: aws.String("test/"),
			}
			listObjectsResp, _ := i.Client.ListObjects(listObjectsParams)
			Expect(*listObjectsResp.Name).To(Equal("image_processing"))
			Expect(*listObjectsResp.Prefix).To(Equal("test/"))
			Expect(*listObjectsResp.Contents[0].Key).To(Equal("test/test.txt"))
			Expect(len(listObjectsResp.Contents)).To(Equal(1))
			getObjectParams := &s3.GetObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String(*listObjectsResp.Contents[0].Key),
			}
			getObjectResp, _ := i.Client.GetObject(getObjectParams)
			defer getObjectResp.Body.Close()
			brb := bytes.Buffer{}
			brb.ReadFrom(getObjectResp.Body)
			srb := brb.String()
			Expect(srb).To(Equal("test!!"))
			// // Delete
			deleteObjectParams := &s3.DeleteObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test/test.txt"),
			}
			_, err = i.Client.DeleteObject(deleteObjectParams)
			Expect(err).To(BeNil())
		})
	})
	Context("NewImageInfra", func() {
		It("Should return Image object", func() {
			i := NewImageMagickInfra(&imageMagickConfig)
			Expect(i.ConvertTo).To(Equal("jpeg"))
		})
	})
})
