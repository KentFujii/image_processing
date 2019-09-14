package config

type hpConfig struct {
	UserAgent string `mapstructure:"user_agent"`
}

func (h *hpConfig) ReadUserAgent() string {
	return h.UserAgent
}
