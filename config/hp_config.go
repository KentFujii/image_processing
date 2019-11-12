package config

type hpConfig struct {
	UserAgent string `mapstructure:"user_agent"`
	Retry int `mapstructure:"retry"`
}

func (c *hpConfig) ReadUserAgent() string {
	return c.UserAgent
}

func (c *hpConfig) ReadRetry() int {
	return c.Retry
}
