package domain

func NewImageMagickInfra(c imageMagickConfig) imageMagickInfra {
	convertTo := c.ReadConvertTo()
	formatWhitelist := c.ReadFormatWhitelist()
	resizeToLimit := c.ReadResizeToLimit()
	resizeToFit := c.ReadResizeToFit()
	return imageMagickInfra{ConvertTo: convertTo, FormatWhitelist: formatWhitelist, ResizeToLimit: resizeToLimit, ResizeToFit: resizeToFit}
}
