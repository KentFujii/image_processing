package infra

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Config interface {
	ReadAwsAccountKey() string
	ReadAwsSecretKey() string
	ReadAwsRegion() string
	ReadAwsEndpoint() string
	ReadBucket() string
}

type imageMagickConfig interface {
	ReadConvertTo() string
	ReadFormatWhitelist() []string
	ReadResizeToLimit() map[string]int
	ReadResizeToFit() map[string]int
}

func NewS3Infra(c s3Config) s3Infra {
	s, _ := session.NewSession()
	aak := c.ReadAwsAccountKey()
	ask := c.ReadAwsSecretKey()
	ar := c.ReadAwsRegion()
	ae := c.ReadAwsEndpoint()
	b := c.ReadBucket()
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials(aak, ask, ""),
		Region: aws.String(ar),
		Endpoint: aws.String(ae),
		S3ForcePathStyle: aws.Bool(true),
	}
	return s3Infra{Client: s3.New(s, &cfg), Bucket: b}
}

func NewImageMagickInfra(c imageMagickConfig) imageMagickInfra {
	convertTo := c.ReadConvertTo()
	formatWhitelist := c.ReadFormatWhitelist()
	resizeToLimit := c.ReadResizeToLimit()
	resizeToFit := c.ReadResizeToFit()
	return imageMagickInfra{ConvertTo: convertTo, FormatWhitelist: formatWhitelist, ResizeToLimit: resizeToLimit, ResizeToFit: resizeToFit}
}
