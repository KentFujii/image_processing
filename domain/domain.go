package domain

// https://medium.com/eureka-engineering/golang-embedded-ac43201cf772
// carrierwaveやimage_supporterの中身を書く
// https://github.com/disintegration/imaging
// ['jpg', 'jpeg', 'gif', 'png', '']
// process resize_to_limit: [600, 600]
// process convert: 'jpg'
// SecureRandom.uuidは使わない
// 画像名のmd5値を使う

func NewImageDomain() imageDomain {
	return imageDomain{}
}
