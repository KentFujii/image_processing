package domain

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDomain(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Domain Suite")
}

type mockImageMagickConfig struct {
	Convert string
	FormatWhitelist []string
	ResizeToLimit map[string]int
	ResizeToFit map[string]int
}

func (c *mockImageMagickConfig) ReadConvertTo() string {
	return c.Convert
}
func (c *mockImageMagickConfig) ReadFormatWhitelist() []string {
	return c.FormatWhitelist
}

func (c *mockImageMagickConfig) ReadResizeToLimit() map[string]int {
	return c.ResizeToLimit
}

func (c *mockImageMagickConfig) ReadResizeToFit() map[string]int {
	return c.ResizeToFit
}



var _ = Describe("Infra", func() {
	var imageMagickConfig mockImageMagickConfig
	BeforeEach(func() {
		imageMagickConfig = mockImageMagickConfig{
			Convert: "jpeg",
			FormatWhitelist: []string{"jpeg", "gif", "png"},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
			ResizeToFit: map[string]int{"height": 100, "width": 100},
		}
	})
	Context("NewImageMagickInfra", func() {
		It("Should return ImageMagick object", func() {
			i := NewImageMagickInfra(&imageMagickConfig)
			Expect(i.ConvertTo).To(Equal("jpg"))
		})
	})
})
