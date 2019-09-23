package infra

import (
	// "fmt"
	"bytes"
	"golang.org/x/xerrors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/common-examples.html
// https://godoc.org/github.com/aws/aws-sdk-go/service/s3/s3iface
// https://aws.amazon.com/jp/blogs/developer/mocking-out-then-aws-sdk-for-go-for-unit-testing/
type s3Infra struct {
	Client s3iface.S3API
	Bucket string
}

func (i *s3Infra) Create(key string, content string) error {
	putObjectParams := &s3.PutObjectInput{
		Bucket: aws.String(i.Bucket),
		Key: aws.String(key),
		Body: bytes.NewReader([]byte(content)),
		ContentType: aws.String("image/jpeg"),
	}
	_, err := i.Client.PutObject(putObjectParams)
	if err != nil {
		return xerrors.Errorf("Create error: %w", err)
	}
	return nil
}

func (i *s3Infra) Read(prefix string) map[string][]byte {
	// https://blog.narumium.net/2019/03/13/%E3%80%90go%E3%80%91aws-s3%E3%81%AB%E5%AF%BE%E3%81%99%E3%82%8B%E5%85%A5%E5%87%BA%E5%8A%9B/
	// https://golang.org/pkg/image/
	// https://godoc.org/github.com/aws/aws-sdk-go/service/s3#GetObjectOutput
	// https://socketloop.com/tutorials/golang-convert-byte-to-image
	// listObjectsParams := &s3.ListObjectsInput{
	// 	Bucket: aws.String(i.Bucket),
	// 	Prefix: aws.String(prefix),
	// }
	// listObjectsResp, _ := i.Client.ListObjects(listObjectsParams)
	// for _, content := range(listObjectsResp.Contents) {
	// 	fmt.Println(*content.Key)
	// }
	// getObjectParams := &s3.GetObjectInput{
	// 	Bucket: aws.String(i.Bucket),
	// 	Key: aws.String("test.txt"),
	// }
	// getObjectResp, _ := i.Client.GetObject(getObjectParams)
	// defer getObjectResp.Body.Close()
	// brb := new(bytes.Buffer)
	// brb.ReadFrom(getObjectResp.Body)
	// srb := brb.Bytes()
	aaa := map[string][]byte{
		prefix: []byte("test!!!"),
	}
	return aaa
}
