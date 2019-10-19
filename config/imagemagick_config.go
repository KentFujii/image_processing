package config

type imageMagickConfig struct {
	Convert string `mapstructure:"convert"`
	ExtensionWhitelist []string `mapstructure:"extension_whitelist"`
	ResizeToLimit map[string]int `mapstructure:"resize_to_limit"`
}

func (c *imageMagickConfig) ReadConvert() string {
	return c.Convert
}

func (c *imageMagickConfig) ReadExtensionWhitelist() []string {
	return c.ExtensionWhitelist
}

func (c *imageMagickConfig) ReadResizeToLimit() map[string]int {
	return c.ResizeToLimit
}
