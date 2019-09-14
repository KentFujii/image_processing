package config

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Infra Suite")
}

var _ = Describe("S3", func() {
	// BeforeEach(func() {
	// 		os.Setenv("GO_ENV", "test")
	// })
	Context("LoadS3Config", func() {
		It("Should create s3 object", func() {
			i := NewS3Infra()
			Expect(i.Create()).To(Equal("image_processing"))
		})
	})
})
