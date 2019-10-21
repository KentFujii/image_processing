package infra

import (
	"fmt"
	"runtime"
	"gopkg.in/gographics/imagick.v2/imagick"
)

var _ = Describe("s3Infra", func() {
	var infra s3Infra
	BeforeEach(func() {
		infra = imageMagickInfra{
			Convert: "jpg",
			FormatWhitelist: []string{"jpeg", "gif", "png"},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
			ResizeToFit: map[string]int{"height": 100, "width": 100},
		}
	})
	Context("ConvertImage", func() {
		It("Should convert image to jpg", func() {
			image, err := os.Open(filePath("testdata/jpg/test0.jpg"))
			infra.ConvertImage()
			// Expect().To(BeNil())
		})
	})
})

func filePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
