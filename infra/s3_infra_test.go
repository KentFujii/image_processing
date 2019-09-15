package infra

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// https://www.google.co.jp/search?{google:acceptedSuggestion}oq=golang+s3iface&{google:instantFieldTrialGroupParameter}sourceid=chrome&ie=UTF-8&q=golang+s3iface
type mockS3Client struct {
    s3iface.S3API
}

var _ = Describe("s3Infra", func() {
	Context("Create", func() {
		It("Should create s3 object", func() {
		})
	})
})
