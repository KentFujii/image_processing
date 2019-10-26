package domain

import (
	"path"
	"runtime"
	"os"
	"fmt"
	"bytes"
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
			ConvertTo: "jpg",
			FormatWhitelist: []string{"jpeg", "gif", "png"},
			ResizeToLimit: map[string]int{"height": 600, "width": 600},
			ResizeToFit: map[string]int{"height": 100, "width": 100},
		}
	})
	Context("ConvertImage", func() {
		It("Should convert png image to jpg", func() {
			file, _ := os.Open(filePath("testdata/png/ocean-1mb.png"))
			defer file.Close()
			brb := bytes.Buffer{}
			brb.ReadFrom(file)
			fmt.Println(domain.ConvertImage(brb.Bytes()))
			Expect(nil).To(BeNil())
		})
	})
})