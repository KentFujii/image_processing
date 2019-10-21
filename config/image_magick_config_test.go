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
			ConvertTo: "jpg",
			FormatWhitelist: []string{"jpg", "jpeg", "gif", "png", ""},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
			ResizeToFit: map[string]int{"height": 100, "width": 100},
		}

	})
	Context("ReadConvert", func() {
		It("Should read convert", func() {
			Expect(config.ReadConvertTo()).To(Equal("jpg"))
		})
	})
	Context("ReadFormatWhitelist", func() {
		It("Should read format whitelist", func() {
			equal := reflect.DeepEqual(config.ReadFormatWhitelist(), []string{"jpg", "jpeg", "gif", "png", ""})
			Expect(equal).To(Equal(true))
		})
	})
	Context("ReadResizeToLimit", func() {
		It("Should read resize to limit", func() {
			equal := reflect.DeepEqual(config.ReadResizeToLimit(), map[string]int{"height": 600, "width": 600})
			Expect(equal).To(Equal(true))
		})
	})
	Context("ReadResizeToFit", func() {
		It("Should read resize to fit", func() {
			equal := reflect.DeepEqual(config.ReadResizeToFit(), map[string]int{"height": 100, "width": 100})
			Expect(equal).To(Equal(true))
		})
	})
})
