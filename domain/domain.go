package domain

type imageConfig interface {
	ReadConvertTo() string
	ReadFormatWhitelist() []string
	ReadResizeToLimit() map[string]int
	ReadResizeToFit() map[string]int
}

func NewImageDomain(c imageConfig) imageDomain {
	convertTo := c.ReadConvertTo()
	formatWhitelist := c.ReadFormatWhitelist()
	resizeToLimit := c.ReadResizeToLimit()
	resizeToFit := c.ReadResizeToFit()
	return imageDomain{ConvertTo: convertTo, FormatWhitelist: formatWhitelist, ResizeToLimit: resizeToLimit, ResizeToFit: resizeToFit}
}
