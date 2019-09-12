package config

import (
	"os"
	"github.com/spf13/viper"
)

// https://www.netlify.com/blog/2016/09/06/creating-a-microservice-boilerplate-in-go/
func LoadS3Config() s3Config {
	env := os.Getenv("GO_ENV")
	viper.SetConfigName(env)
	viper.AddConfigPath("/go/src/config/env/")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	c := s3Config{}
	viper.UnmarshalKey("s3", &c)
	return c
}

func LoadHpConfig() hpConfig {
	env := os.Getenv("GO_ENV")
	viper.SetConfigName(env)
	viper.AddConfigPath("/go/src/config/env/")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()
	h := hpConfig{}
	viper.UnmarshalKey("hp", &h)
	return h
}
