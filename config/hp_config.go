package config

type hpConfig struct {
	UserAgent string `mapstructure:"user_agent"`
}

func (h *hpConfig) fetchUserAgent() string {
	return h.UserAgent
}
