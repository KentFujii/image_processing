package domain

type imageConfig interface {
	ReadConvertTo() string
	ReadFormatWhitelist() []string
	ReadResizeToLimit() map[string]int
	ReadResizeToFit() map[string]int
}

func NewImageDomain(c imageConfig) imageMagickInfra {
	convertTo := c.ReadConvertTo()
	formatWhitelist := c.ReadFormatWhitelist()
	resizeToLimit := c.ReadResizeToLimit()
	resizeToFit := c.ReadResizeToFit()
	return imageMagickInfra{ConvertTo: convertTo, FormatWhitelist: formatWhitelist, ResizeToLimit: resizeToLimit, ResizeToFit: resizeToFit}
}
