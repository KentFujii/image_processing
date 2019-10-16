package config

type s3Config struct {
	AwsAccountKey string `mapstructure:"aws_account_key"`
	AwsSecretKey string `mapstructure:"aws_secret_key"`
	AwsRegion string `mapstructure:"aws_region"`
	AwsEndpoint string `mapstructure:"aws_endpoint"`
	Bucket string `mapstructure:"bucket"`
}

func (c *s3Config) ReadAwsAccountKey() string {
	return c.AwsAccountKey
}

func (c *s3Config) ReadAwsSecretKey() string {
	return c.AwsSecretKey
}

func (c *s3Config) ReadAwsRegion() string {
	return c.AwsRegion
}

func (c *s3Config) ReadAwsEndpoint() string {
	return c.AwsEndpoint
}

func (c *s3Config) ReadBucket() string {
	return c.Bucket
}
