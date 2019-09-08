package config

import (
	"os"
	"github.com/spf13/viper"
)

type S3Config struct {
	AwsAccountKey string `mapstructure:"aws_account_key"`
	AwsSecretKey string `mapstructure:"aws_secret_key"`
	AwsRegion string `mapstructure:"aws_region"`
	AwsEndpoint string `mapstructure:"aws_endpoint"`
	Bucket string `mapstructure:"bucket"`
}

type HPConfig struct {
	UserAgent string `mapstructure:"user_agent"`
}


type Config struct {
	S3 S3Config `mapstructure:"s3"`
	HP HPConfig `mapstructure:"hp"`
}

func SetUp() Config {
	env := os.Getenv("GO_ENV")
	viper.SetConfigName(env)
	viper.AddConfigPath("/go/src/config/")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	c := Config{}
	viper.Unmarshal(&c)
	return c
}
