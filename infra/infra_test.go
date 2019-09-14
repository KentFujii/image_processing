package infra

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
	Context("LoadS3Config", func() {
		It("Should create s3 object", func() {
			i := NewS3Infra()
			i.ReadObjects()
		})
	})
})
