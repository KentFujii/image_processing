package infra

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
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
			fmt.Println(infra)
			// Expect(infra.Create()).To(Equal("hoge"))
		})
	})
})
