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

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infra Suite")
}

var _ = Describe("Infra", func() {
	// ['jpg', 'jpeg', 'gif', 'png', '']
	// process resize_to_limit: [600, 600]
	// process convert: 'jpg'
	// SecureRandom.uuidは使わない
	// 画像名のmd5値を使う
	// infra層では上記を扱わない
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
		It("Should crud s3 object", func() {
			// テストデータからtestdataからのjpg使え
			i := NewS3Infra(&s3Config)
			var putObjectParams *s3.PutObjectInput
			var err error
			// create
			putObjectParams = &s3.PutObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test.txt"),
				Body: bytes.NewReader([]byte("test!")),
				ContentType: aws.String("text/plain"),
			}
			_, err = i.Client.PutObject(putObjectParams)
			Expect(err).To(BeNil())
			// read
			listObjectsParams := &s3.ListObjectsInput{
				Bucket: aws.String(i.Bucket),
				Prefix: aws.String("test.txt"),
			}
			listObjectsResp, _ := i.Client.ListObjects(listObjectsParams)
			Expect(len(listObjectsResp.Contents)).To(Equal(1))
			getObjectParams := &s3.GetObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test.txt"),
			}
			getObjectResp, _ := i.Client.GetObject(getObjectParams)
			defer getObjectResp.Body.Close()
			brb := new(bytes.Buffer)
			brb.ReadFrom(getObjectResp.Body)
			srb := brb.String()
			Expect(srb).To(Equal("test!"))
			// update
			putObjectParams = &s3.PutObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test.txt"),
				Body: bytes.NewReader([]byte("test!!")),
				ContentType: aws.String("text/plain"),
			}
			_, err = i.Client.PutObject(putObjectParams)
			Expect(err).To(BeNil())
			// delete
			deleteObjectParams := &s3.DeleteObjectInput{
				Bucket: aws.String(i.Bucket),
				Key: aws.String("test.txt"),
			}
			_, err = i.Client.DeleteObject(deleteObjectParams)
			Expect(err).To(BeNil())
		})
	})
})
