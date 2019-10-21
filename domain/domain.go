package domain

type imageMagickConfig interface {
	ReadConvertTo() string
	ReadFormatWhitelist() []string
	ReadResizeToLimit() map[string]int
	ReadResizeToFit() map[string]int
}

func NewImageDomain(c imageMagickConfig) imageMagickInfra {
	convertTo := c.ReadConvertTo()
	formatWhitelist := c.ReadFormatWhitelist()
	resizeToLimit := c.ReadResizeToLimit()
	resizeToFit := c.ReadResizeToFit()
	return imageMagickInfra{ConvertTo: convertTo, FormatWhitelist: formatWhitelist, ResizeToLimit: resizeToLimit, ResizeToFit: resizeToFit}
}
