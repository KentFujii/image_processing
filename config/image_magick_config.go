package config

type imageConfig struct {
	ConvertTo string `mapstructure:"convert_to"`
	FormatWhitelist []string `mapstructure:"format_whitelist"`
	ResizeToLimit map[string]int `mapstructure:"resize_to_limit"`
	ResizeToFit map[string]int `mapstructure:"resize_to_fit"`
}

func (c *imageConfig) ReadConvertTo() string {
	return c.ConvertTo
}

func (c *imageConfig) ReadFormatWhitelist() []string {
	return c.FormatWhitelist
}

func (c *imageConfig) ReadResizeToLimit() map[string]int {
	return c.ResizeToLimit
}

func (c *imageConfig) ReadResizeToFit() map[string]int {
	return c.ResizeToFit
}
