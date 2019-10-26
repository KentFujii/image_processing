package domain

import (
	"path"
	"runtime"
	"os"
	"bytes"
	"image"
	_ "image/jpeg"
	_ "image/gif"
	_ "image/png"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func filePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}

var _ = Describe("s3Infra", func() {
	var domain imageDomain
	BeforeEach(func() {
		domain = imageDomain{
			ConvertTo: "jpeg",
			FormatWhitelist: []string{"jpeg", "gif", "png"},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
			ResizeToFit: map[string]int{"height": 100, "width": 100},
		}
	})
	// https://www.admfactory.com/how-to-get-the-dimensions-of-an-image-in-golang/
	// https://gist.github.com/akhenakh/8462840
	Context("ConvertFormat", func() {
		It("Should convert jpeg image to jpeg", func() {
			file, _ := os.Open(filePath("testdata/jpeg/airplane-1mb.jpg"))
			defer file.Close()
			inputBrb := bytes.Buffer{}
			inputBrb.ReadFrom(file)
			inputBin := inputBrb.Bytes()
			outputBin, _ := domain.ConvertFormat(inputBin)
			outputBrb := bytes.NewReader(outputBin)
			_, format, _ := image.DecodeConfig(outputBrb)
			Expect(format).To(Equal("jpeg"))
		})
		It("Should convert png image to jpeg", func() {
			file, _ := os.Open(filePath("testdata/png/ocean-1mb.png"))
			defer file.Close()
			inputBrb := bytes.Buffer{}
			inputBrb.ReadFrom(file)
			inputBin := inputBrb.Bytes()
			outputBin, _ := domain.ConvertFormat(inputBin)
			outputBrb := bytes.NewReader(outputBin)
			_, format, _ := image.DecodeConfig(outputBrb)
			Expect(format).To(Equal("jpeg"))
		})
	})
	// Context("ResizeImageToLimit", func() {
	// 	It("Should resize image to limit", func() {
	// 		file, _ := os.Open(filePath("testdata/jpeg/airplane-1mb.jpg"))
	// 		defer file.Close()
	// 		inputBrb := bytes.Buffer{}
	// 		inputBrb.ReadFrom(file)
	// 		inputBin := inputBrb.Bytes()
	// 		outputBin, _ := domain.ResizeImageToLimit(inputBin)
	// 		Expect(format).To(Equal("image/jpeg"))
	// 	})
	// })
})
