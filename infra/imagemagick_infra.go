package infra

// https://github.com/gographics/imagick
// https://godoc.org/gopkg.in/gographics/imagick.v2/imagick
// https://github.com/gographics/imagick/blob/master/examples/docker/main.go

import (
	"gopkg.in/gographics/imagick.v2/imagick"
)

type imageMagickInfra struct {
	Client imagick
}
