package domain

type s3Infra interface {
	Put(key string, content string, contentType string) error
	List(prefix string) ([]string, error)
	Get(key string) ([]byte, error)
	Delete(key string) error
}

type imageDomain struct {}

// // 対象元の画像ファイルをjpgとしてE/transform
// func (d *s3Domain) PullImagesFromS3(prefix string, contentType string) error {
// }

// // Convert
// https://github.com/gographics/imagick
// func (d *s3Domain) ConvertImage(prefix string, contentType string) error {
// }

// // Validate
// func (d *s3Domain) ValidateImage(prefix string, contentType string) error {
// }

// // 対象先に画像ファイルをjpgとしてvalidation/store
// func (d *s3Domain) PushImagesToS3(key string, content string, contentType string) error {
// }
