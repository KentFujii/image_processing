package config

type hpConfig struct {
	UserAgent string `mapstructure:"user_agent"`
}

func (c *hpConfig) ReadUserAgent() string {
	return c.UserAgent
}
