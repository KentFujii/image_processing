package infra

import (
	"bytes"
	"golang.org/x/xerrors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

type s3Infra struct {
	Client s3iface.S3API
	Bucket string
}

func (i *s3Infra) Put(key string, content string, contentType string) error {
	putObjectParams := &s3.PutObjectInput{
		Bucket: aws.String(i.Bucket),
		Key: aws.String(key),
		Body: bytes.NewReader([]byte(content)),
		ContentType: aws.String(contentType),
	}
	_, err := i.Client.PutObject(putObjectParams)
	if err != nil {
		return xerrors.Errorf("Create error: %w", err)
	}
	return nil
}

// ListObjectsとGetObjectを分離する
func (i *s3Infra) List(prefix string) map[string][]byte {
	listObjectsParams := &s3.ListObjectsInput{
		Bucket: aws.String(i.Bucket),
		Prefix: aws.String(prefix),
	}
	listObjectsResp, _ := i.Client.ListObjects(listObjectsParams)
	bodies := map[string][]byte{}
	var getObjectParams *s3.GetObjectInput
	var getObjectResp *s3.GetObjectOutput
	var brb bytes.Buffer
	for _, content := range(listObjectsResp.Contents) {
		getObjectParams = &s3.GetObjectInput{
			Bucket: aws.String(i.Bucket),
			Key: aws.String(*content.Key),
		}
		getObjectResp, _ = i.Client.GetObject(getObjectParams)
		defer getObjectResp.Body.Close()
		brb = bytes.Buffer{}
		brb.ReadFrom(getObjectResp.Body)
		bodies[*content.Key] = brb.Bytes()
	}
	return bodies
}

func (i *s3Infra) Delete(key string) error {
	deleteObjectParams := &s3.DeleteObjectInput{
		Bucket: aws.String(i.Bucket),
		Key: aws.String(key),
	}
	_, err := i.Client.DeleteObject(deleteObjectParams)
	if err != nil {
		return xerrors.Errorf("Delete error: %w", err)
	}
	return nil
}
