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
			c := NewS3Config()
			Expect(c.ReadAwsAccountKey()).To(Equal("image_processing"))
			Expect(c.ReadAwsSecretKey()).To(Equal("password"))
			Expect(c.ReadAwsRegion()).To(Equal("ap-northeast-1"))
			Expect(c.ReadAwsEndpoint()).To(Equal("http://storage:9000"))
			Expect(c.ReadBucket()).To(Equal("image_processing"))
		})
	})

	Context("LoadHpConfig", func() {
		It("should load hp config", func() {
			c := NewHpConfig()
			Expect(c.ReadUserAgent()).To(Equal("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/602.4.8 (KHTML, like Gecko) Version/10.0.3 Safari/602.4.8"))
		})
	})
})
