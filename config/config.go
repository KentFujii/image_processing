package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var DefaultConfig Config

type S3Config struct {
	Host string `mapstructure:"host"`
	ImageBucket string `mapstructure:"image_bucket"`
}

type HPConfig struct {
	UserAgent string `mapstructure:"user_agent"`
}


type Config struct {
	S3 S3Config `mapstructure:"s3"`
	HP HPConfig `mapstructure:"hp"`
}

func Setup() {
	flag()
	yaml()
}

func flag() {
	pflag.String("s", "local", "local/dev/stg/prd")
	pflag.Int("worker", 3, "worker int value")
	pflag.Int("queue", 10000, "queue int value")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	log.Printf("loaded pflag")
}

func yaml() error {
	viper.SetConfigName(viper.GetString("s"))
	viper.SetConfigType("yaml")
	for _, path := range []string{"./", "/go/bin/config"} {
		viper.AddConfigPath(path)
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		return err
	}
	if err := viper.Unmarshal(&DefaultConfig); err != nil {
		panic(err)
	}

	fmt.Println(DefaultConfig.HP.UserAgent)
	return nil
}
