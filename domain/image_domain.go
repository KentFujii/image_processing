package domain

import (
	"fmt"
	"bytes"
	"os/exec"
)

type imageDomain struct {
	ConvertTo string
	FormatWhitelist []string
	ResizeToLimit map[string]int
	ResizeToFit map[string]int
}

// https://socketloop.com/tutorials/golang-convert-an-image-file-to-byte
// https://socketloop.com/tutorials/golang-convert-byte-to-image
// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/functions/imagemagick/imagemagick.go
// cat domain/testdata/png/ocean-1mb.png | convert - jpeg:- | identify -
func (i *imageDomain) ConvertFormat(bin []byte) []byte {
	// Read
	// Convert
	// Resize
	input := bytes.NewReader(bin)
	var output bytes.Buffer
	cmd := exec.Command("convert", "-")
	cmd.Stdin = input
	cmd.Stdout = &output
	cmd.Run()
	fmt.Println(output.String())
	return bin
}

// func (i *imageMagickInfra) CompareImage(sourceBlob []byte, targetBlob []byte) bool {
// 	// magick_local_image = Magick::Image.from_blob(local_image_bin).first
// 	// magick_remote_image = Magick::Image.from_blob(remote_image_bin).first
// 	// local_small_image = magick_local_image.resize_to_fit(100)
// 	// remote_small_image = magick_remote_image.resize_to_fit(100)
// 	// diff = local_small_image.composite(remote_small_image, 0, 0, Magick::DifferenceCompositeOp)
// 	// diff.channel_mean.first.to_i <= 3500
// 	return true
// }
