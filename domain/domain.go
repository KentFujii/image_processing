package domain

type s3Infra interface {
	Put(key string, content string, contentType string) error
	List(prefix string) ([]string, error)
	Get(key string) ([]byte, error)
	Delete(key string) error
}

func NewImageDomain(i s3Infra) s3Domain {
	return imageDomain{S3: s3Infra}
}
