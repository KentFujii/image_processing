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

var _ = Describe("Domain", func() {
	// var s3Config mockS3Config
	// BeforeEach(func() {
	// })
	// Context("NewS3Domain", func() {
	// })
})
