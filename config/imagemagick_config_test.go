package config

import (
	"reflect"
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
	Context("ReadExtensionWhitelist", func() {
		It("Should read extension whitelist", func() {
			equal := reflect.DeepEqual(config.ReadExtensionWhitelist(), []string{"jpg", "jpeg", "gif", "png", ""})
			Expect(equal).To(Equal(true))
		})
	})
	Context("ReadResizeToLimit", func() {
		It("Should read resize to limit", func() {
			equal := reflect.DeepEqual(config.ReadResizeToLimit(), map[string]int{"height": 600, "width": 600})
			Expect(equal).To(Equal(true))
		})
	})
})
