package config

import (
	"os"
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("config", func() {
	BeforeEach(func() {
			os.Setenv("GO_ENV", "test")
	})

	Context("SetUp", func() {
		It("should setup config", func() {
			setUp := SetUp()
			Expect(setUp.S3.AwsAccountKey).To(Equal("image_processing"))
		})
	})
})
