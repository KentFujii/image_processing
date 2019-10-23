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

type mockImageConfig struct {
	ConvertTo string
	FormatWhitelist []string
	ResizeToLimit map[string]int
	ResizeToFit map[string]int
}

func (c *mockImageConfig) ReadConvertTo() string {
	return c.ConvertTo
}
func (c *mockImageConfig) ReadFormatWhitelist() []string {
	return c.FormatWhitelist
}

func (c *mockImageConfig) ReadResizeToLimit() map[string]int {
	return c.ResizeToLimit
}

func (c *mockImageConfig) ReadResizeToFit() map[string]int {
	return c.ResizeToFit
}

var _ = Describe("Infra", func() {
	var imageConfig mockImageConfig
	BeforeEach(func() {
		imageConfig = mockImageConfig{
			ConvertTo: "jpeg",
			FormatWhitelist: []string{"jpeg", "gif", "png"},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
			ResizeToFit: map[string]int{"height": 100, "width": 100},
		}
	})
	Context("NewImageInfra", func() {
		It("Should return Image object", func() {
			i := NewImageDomain(&imageConfig)
			Expect(i.ConvertTo).To(Equal("jpeg"))
		})
	})
})
