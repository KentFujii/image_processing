package infra

import (
	"bytes"
	"io/ioutil"
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

func (m *mockS3Client) ListObjects(input *s3.ListObjectsInput) (*s3.ListObjectsOutput, error) {
	key0 := "test/test0.txt"
	key1 := "test/test1.txt"
	object0 := &s3.Object{
		Key: &key0,
	}
	object1 := &s3.Object{
		Key: &key1,
	}
	objects := []*s3.Object{object0, object1}
	return &s3.ListObjectsOutput{
		Contents: objects,
	}, nil
}


func (m *mockS3Client) GetObject(input *s3.GetObjectInput) (*s3.GetObjectOutput, error) {
	body := ioutil.NopCloser(bytes.NewReader([]byte("test body!")))
	return &s3.GetObjectOutput{
		Body: body,
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
			Expect(infra.Create("testKey", "testContent", "image/jpeg")).To(BeNil())
		})
	})
	Context("Read", func() {
		It("Should read s3 object", func() {
			bodies := infra.Read("testPrefix")
			Expect(len(bodies)).To(Equal(2))
		})
	})
})
