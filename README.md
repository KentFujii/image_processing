https://github.com/hmarf/sample_LayerArchitecture

https://disintegration.dev/

https://rooter.jp/aws/aws-sdk-go-s3-four-tips/

https://qiita.com/Azizim_A/items/66564b5dc7597717932b#dep%E3%81%AE%E4%BD%BF%E3%81%84%E6%96%B9

http://onsi.github.io/ginkgo/

https://github.com/golang/go/wiki/Modules#quick-start

https://www.netlify.com/blog/2016/09/06/creating-a-microservice-boilerplate-in-go/

https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/s3

https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/s3/create_new_bucket_and_object.go

https://aws.amazon.com/jp/blogs/developer/mocking-out-then-aws-sdk-for-go-for-unit-testing/

https://dev.classmethod.jp/go/access-minio-using-aws-sdk-for-go/

https://qiita.com/hmarf/items/7f4d39c48775c205b99b

https://qiita.com/dich1/items/0871c9c1ff058d0d693b

```
ui
app
domain
infra
config
```

- ui
  - a layer manipulated from outer commands
  - ex. a cli of batch or a view of MVC
- app
  - a layer scripting a behavior of the application
  - ex. a controller of MVC
- domain
  - a layer defining the business logic
  - ex. a model of MVC
- infra
  - a layer wrapping IO
  - ex. interacting with IO like ORM
