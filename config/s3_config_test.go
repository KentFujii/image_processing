package config

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("s3Config", func() {
	var config s3Config
	BeforeEach(func() {
		config = s3Config{
			AwsAccountKey: "image_processing",
			AwsSecretKey: "password",
			AwsRegion: "ap-northeast-1",
			AwsEndpoint: "http://storage:9000",
			Bucket: "image_processing",
		}

	})
	Context("ReadAwsAccountKey", func() {
		It("Should read aws account key", func() {
			Expect(config.ReadAwsAccountKey()).To(Equal("image_processing"))
		})
	})
	Context("ReadAwsSecretKey", func() {
		It("Should read aws secret key", func() {
			Expect(config.ReadAwsSecretKey()).To(Equal("password"))
		})
	})
	Context("ReadAwsRegion", func() {
		It("Should read aws region", func() {
			Expect(config.ReadAwsRegion()).To(Equal("ap-northeast-1"))
		})
	})
	Context("ReadAwsEndpoint", func() {
		It("Should read aws endpoint", func() {
			Expect(config.ReadAwsEndpoint()).To(Equal("http://storage:9000"))
		})
	})
	Context("ReadBucket", func() {
		It("Should read bucket", func() {
			Expect(config.ReadBucket()).To(Equal("image_processing"))
		})
	})
})
