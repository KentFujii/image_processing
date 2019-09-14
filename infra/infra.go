package infra

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3Infra() s3Infra {
}

// type s3 struct {
// }

// type hp struct {
// }

// func (s *S3) fetch(id int) (err error) {
// 	post.Id = id
// 	return
// }

// func (post *S3) create() (err error) {
// 	return
// }

// func (post *S3) update() (err error) {
// 	return
// }

// func (post *S3) delete() (err error) {
// 	return
// }

// https://dev.classmethod.jp/go/access-minio-using-aws-sdk-for-go/
// https://qiita.com/hmarf/items/7f4d39c48775c205b99b
// func newS3() (*s3.S3, error) {
//   s, err := session.NewSession()
//   if err != nil {
//     return nil, err
//   }

//   ak := "image_processing"
//   sk := "password"
//   cfg := aws.Config{
//     Credentials: credentials.NewStaticCredentials(ak, sk, ""),
//     Region:      aws.String("ap-northeast-1"),
//     Endpoint:    aws.String("http://storage:9000"),
//     S3ForcePathStyle: aws.Bool(true),
//   }
//   return s3.New(s, &cfg), nil
// }

// func Transfer() {
//   c, err := newS3()
//   if err != nil {
//     panic(err)
//   }

//   bucket := "image_processing"
//   var token *string
//   for complete := false; !complete; {
//     in := s3.ListObjectsV2Input{Bucket: &bucket, ContinuationToken: token}
//     out, err := c.ListObjectsV2(&in)
//     if err != nil {
//       panic(err)
//     }

//     for i, o := range out.Contents {
//       fmt.Printf("[%d] : %s\n", i, *o.Key)
//     }

//     complete = out.IsTruncated != nil && !*out.IsTruncated
//     token = out.NextContinuationToken
//   }
// }
