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

func (i *s3Infra) Put(key string, content []byte, contentType string) error {
	putObjectParams := &s3.PutObjectInput{
		Bucket: aws.String(i.Bucket),
		Key: aws.String(key),
		Body: bytes.NewReader(content),
		ContentType: aws.String(contentType),
	}
	_, err := i.Client.PutObject(putObjectParams)
	if err != nil {
		return xerrors.Errorf("Create error: %w", err)
	}
	return nil
}

func (i *s3Infra) List(prefix string) ([]string, error) {
	listObjectsParams := &s3.ListObjectsInput{
		Bucket: aws.String(i.Bucket),
		Prefix: aws.String(prefix),
	}
	listObjectsResp, err := i.Client.ListObjects(listObjectsParams)
	if err != nil {
		return nil, xerrors.Errorf("List error: %w", err)
	}
	keys := []string{}
	for _, content := range(listObjectsResp.Contents) {
		keys = append(keys, *content.Key)
	}
	return keys, nil
}

func (i *s3Infra) Get(key string) ([]byte, error) {
	getObjectInput := &s3.GetObjectInput{
		Bucket: aws.String(i.Bucket),
		Key: aws.String(key),
	}
	getObject, err := i.Client.GetObject(getObjectInput)
	if err != nil {
		return nil, xerrors.Errorf("Get error: %w", err)
	}
	defer getObject.Body.Close()
	brb := bytes.Buffer{}
	brb.ReadFrom(getObject.Body)
	body := brb.Bytes()
	return body, nil
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
