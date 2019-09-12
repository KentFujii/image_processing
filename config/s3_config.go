package config

type s3Config struct {
	AwsAccountKey string `mapstructure:"aws_account_key"`
	AwsSecretKey string `mapstructure:"aws_secret_key"`
	AwsRegion string `mapstructure:"aws_region"`
	AwsEndpoint string `mapstructure:"aws_endpoint"`
	Bucket string `mapstructure:"bucket"`
}

func (s *s3Config) fetchAwsAccountKey() string {
	return s.AwsAccountKey
}

func (s *s3Config) fetchAwsSecretKey() string {
	return s.AwsSecretKey
}

func (s *s3Config) fetchAwsRegion() string {
	return s.AwsRegion
}

func (s *s3Config) fetchAwsEndpoint() string {
	return s.AwsEndpoint
}

func (s *s3Config) fetchBucket() string {
	return s.Bucket
}
