package config

type imageMagickConfig struct {
	ConvertTo string `mapstructure:"convert_to"`
	FormatWhitelist []string `mapstructure:"format_whitelist"`
	ResizeToLimit map[string]int `mapstructure:"resize_to_limit"`
	ResizeToFit map[string]int `mapstructure:"resize_to_fit"`
}

func (c *imageMagickConfig) ReadConvertTo() string {
	return c.ConvertTo
}

func (c *imageMagickConfig) ReadFormatWhitelist() []string {
	return c.FormatWhitelist
}

func (c *imageMagickConfig) ReadResizeToLimit() map[string]int {
	return c.ResizeToLimit
}

func (c *imageMagickConfig) ReadResizeToFit() map[string]int {
	return c.ResizeToFit
}
