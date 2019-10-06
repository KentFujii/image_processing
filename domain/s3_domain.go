package domain
// carrierwaveやimage_supporterの中身を書く
// https://github.com/disintegration/imaging
// ['jpg', 'jpeg', 'gif', 'png', '']
// process resize_to_limit: [600, 600]
// process convert: 'jpg'
// SecureRandom.uuidは使わない
// 画像名のmd5値を使う

type imageDomain struct {
	S3 s3Infra
	Hp hpInfra
}

// 対象元の画像ファイルをjpgとしてE/transform
func (d *s3Domain) PullImagesFromS3(prefix string, contentType string) error {
}

// Convert
func (d *s3Domain) ConvertImage(prefix string, contentType string) error {
}

// Validate
func (d *s3Domain) ValidateImage(prefix string, contentType string) error {
}

// 対象先に画像ファイルをjpgとしてvalidation/store
func (d *s3Domain) PushImagesToS3(key string, content string, contentType string) error {
}
