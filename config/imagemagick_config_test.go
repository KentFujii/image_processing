package config

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("imageMagickConfig", func() {
	var config imageMagickConfig
	BeforeEach(func() {
		config = imageMagickConfig{
			Convert: "jpg",
			ExtensionWhitelist: []string{"jpg", "jpeg", "gif", "png", ""},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
		}

	})
	Context("ReadConvert", func() {
		It("Should read convert", func() {
			Expect(config.ReadConvert()).To(Equal("jpg"))
		})
	})
	// Context("ReadExtensionWhitelist", func() {
	// 	It("Should read extension whitelist", func() {
	// 		Expect(config.ReadExtensionWhitelist()).To(Equal(["jpg", "jpeg", "gif", "png", ""]))
	// 	})
	// })
	// Context("ReadResizeToLimit", func() {
	// 	It("Should read resize to limit", func() {
	// 		Expect(config.ReadResizeToLimit()).To(Equal({"height": 600, "width": 600}))
	// 	})
	// })
})
