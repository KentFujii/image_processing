package domain

import (
	"path"
	"runtime"
	"os"
	"bytes"
	"image"
	// "io"
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

var _ = Describe("imageDomain", func() {
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
	Context("ResizeImageToLimit", func() {
		It("Should resize 689x689 image to limit", func() {
			file, _ := os.Open(filePath("testdata/jpeg/butterfly-100kb.jpg"))
			defer file.Close()
			inputBrb := bytes.Buffer{}
			inputBrb.ReadFrom(file)
			inputBin := inputBrb.Bytes()
			outputBin, _ := domain.ResizeImageToLimit(inputBin)
			outputBrb := bytes.NewReader(outputBin)
			config, _, _ := image.DecodeConfig(outputBrb)
			Expect(config.Height).To(Equal(600))
			Expect(config.Width).To(Equal(600))
		})
		It("Should resize 272x170 image to limit", func() {
			file, _ := os.Open(filePath("testdata/png/mountain-100kb.png"))
			defer file.Close()
			inputBrb := bytes.Buffer{}
			inputBrb.ReadFrom(file)
			inputBin := inputBrb.Bytes()
			outputBin, _ := domain.ResizeImageToLimit(inputBin)
			outputBrb := bytes.NewReader(outputBin)
			config, _, _ := image.DecodeConfig(outputBrb)
			Expect(config.Height).To(Equal(170))
			Expect(config.Width).To(Equal(272))
		})
	})
	Context("CompareImage", func() {
		It("Should compare jpeg files and return true", func() {
			srcFile, _ := os.Open(filePath("testdata/jpeg/butterfly-100kb.jpg"))
			defer srcFile.Close()
			inputSrcBrb := bytes.Buffer{}
			inputSrcBrb.ReadFrom(srcFile)
			inputSrcBin := inputSrcBrb.Bytes()
			dstFile, _ := os.Open(filePath("testdata/jpeg/butterfly-500kb.jpg"))
			defer dstFile.Close()
			inputDstBrb := bytes.Buffer{}
			inputDstBrb.ReadFrom(dstFile)
			inputDstBin := inputDstBrb.Bytes()
			result, _ := domain.CompareImage(inputSrcBin, inputDstBin)
			Expect(result).To(Equal(true))
		})
	})
})
