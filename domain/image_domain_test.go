package domain

import (
	"path"
	"runtime"
	"os"
	"bytes"
	"net/http"
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
	Context("ConvertFormat", func() {
		It("Should convert png image to jpeg", func() {
			file, _ := os.Open(filePath("testdata/png/ocean-1mb.png"))
			defer file.Close()
			inputBrb := bytes.Buffer{}
			inputBrb.ReadFrom(file)
			inputBin := inputBrb.Bytes()
			outputBin, _ := domain.ConvertFormat(inputBin)
			format := http.DetectContentType(outputBin)
			Expect(format).To(Equal("image/jpeg"))
		})
		It("Should convert jpeg image to jpeg", func() {
			file, _ := os.Open(filePath("testdata/jpeg/airplane-1mb.jpg"))
			defer file.Close()
			inputBrb := bytes.Buffer{}
			inputBrb.ReadFrom(file)
			inputBin := inputBrb.Bytes()
			outputBin, _ := domain.ConvertFormat(inputBin)
			format := http.DetectContentType(outputBin)
			Expect(format).To(Equal("image/jpeg"))
		})
	})
	// Context("Resize", func() {
	// 	It("Should resize image to limit", func() {
	// 		file, _ := os.Open(filePath("testdata/png/ocean-1mb.png"))
	// 		defer file.Close()
	// 		inputBrb := bytes.Buffer{}
	// 		inputBrb.ReadFrom(file)
	// 		inputBin := inputBrb.Bytes()
	// 		outputBin, _ := domain.ResizeToLimit(inputBin)
	// 	})
	// })
})
