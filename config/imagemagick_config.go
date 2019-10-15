package config

// https://github.com/gographics/imagick
type imageMagickConfig struct {
	Convert string `mapstructure:"convert"`
	ResizeToLimit map[string]int `mapstructure:"resize_to_limit"`
	ExtensionWhitelist []string `mapstructure:"extension_whitelist"`
}
