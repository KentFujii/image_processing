package config

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("hpConfig", func() {
	var config hpConfig
	BeforeEach(func() {
		config = hpConfig{
			UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/602.4.8 (KHTML, like Gecko) Version/10.0.3 Safari/602.4.8",
			Retry: 3,
		}

	})
	Context("ReadUserAgent", func() {
		It("Should read user agent", func() {
			Expect(config.ReadUserAgent()).To(Equal("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_3) AppleWebKit/602.4.8 (KHTML, like Gecko) Version/10.0.3 Safari/602.4.8"))
		})
	})
	Context("ReadRetry", func() {
		It("Should read retry", func() {
			Expect(config.ReadRetry()).To(Equal(3))
		})
	})

})
