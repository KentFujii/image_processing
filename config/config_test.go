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

	Context("LoadS3Config", func() {
		It("should load s3 config", func() {
			c := LoadS3Config()
			Expect(c.fetchAwsAccountKey()).To(Equal("image_processing"))
			Expect(c.fetchAwsSecretKey()).To(Equal("password"))
			Expect(c.fetchAwsRegion()).To(Equal("ap-northeast-1"))
			Expect(c.fetchAwsEndpoint()).To(Equal("http://storage:9000"))
			Expect(c.fetchBucket()).To(Equal("image_processing"))
		})
	})

	Context("LoadHpConfig", func() {
		It("should load hp config", func() {
			c := LoadHpConfig()
			Expect(c.fetchUserAgent()).To(Equal("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/602.4.8 (KHTML, like Gecko) Version/10.0.3 Safari/602.4.8"))
		})
	})
})
