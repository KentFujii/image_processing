package config

type s3Config struct {
	AwsAccountKey string `mapstructure:"aws_account_key"`
	AwsSecretKey string `mapstructure:"aws_secret_key"`
	AwsRegion string `mapstructure:"aws_region"`
	AwsEndpoint string `mapstructure:"aws_endpoint"`
	Bucket string `mapstructure:"bucket"`
}

func (s *s3Config) ReadAwsAccountKey() string {
	return s.AwsAccountKey
}

func (s *s3Config) ReadAwsSecretKey() string {
	return s.AwsSecretKey
}

func (s *s3Config) ReadAwsRegion() string {
	return s.AwsRegion
}

func (s *s3Config) ReadAwsEndpoint() string {
	return s.AwsEndpoint
}

func (s *s3Config) ReadBucket() string {
	return s.Bucket
}
