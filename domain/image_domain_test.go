package domain

type mockS3Infra struct {
	Client s3iface.S3API
	Bucket string
}

func (i *mockS3Infra) Put(key string, content string, contentType string) error {
}

// func (i *mockS3Infra) List(prefix string) ([]string, error) {
// }

// func (i *mockS3Infra) Get(key string) ([]byte, error) {
// }

// func (i *mockS3Infra) Delete(key string) error {
// }

var _ = Describe("Domain", func() {
	var domain imageDomain
	var s3Infra mockS3Infra
	BeforeEach(func() {
		domain = imageDomain{}
	})
	Context("PullImagesFromS3", func() {
		fmt.Println(domain)
	})
})
