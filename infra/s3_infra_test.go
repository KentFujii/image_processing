package infra

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type mockS3Client struct {
    s3iface.S3API
}

func (m *mockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	etag := "test ETag!"
	return &s3.PutObjectOutput{
		ETag: &etag,
	}, nil
}

// func (m *mockS3Client) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	// return &s3.ListObjectsOutput{
	// 	Contents: ,
	// }, nil
// }

var _ = Describe("s3Infra", func() {
	var infra s3Infra
	BeforeEach(func() {
		infra = s3Infra{
			Client: &mockS3Client{},
			Bucket: "image_processing",
		}

	})
	Context("Create", func() {
		It("Should create s3 object", func() {
			Expect(infra.Create("testKey", "testContent")).To(BeNil())
		})
	})
	Context("Read", func() {
		It("Should read s3 object", func() {
			r := infra.Read("testPrefix")
			Expect(string(r["testPrefix"])).To(Equal("test!!!"))
		})
	})
})
