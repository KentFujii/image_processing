package domain

import (
	"bytes"
	"os/exec"
	"strconv"
	"fmt"
	"golang.org/x/xerrors"
)

type imageDomain struct {
	ConvertTo string
	FormatWhitelist []string
	ResizeToLimit map[string]int
	ResizeToFit map[string]int
}

func (d *imageDomain) ConvertFormat(bin []byte) ([]byte, error) {
	input := bytes.NewReader(bin)
	var output bytes.Buffer
	cmd := exec.Command("convert", "-", d.ConvertTo + ":-")
	cmd.Stdin = input
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		return nil, xerrors.Errorf("ConvertFormat error: %w", err)
	}
	return output.Bytes(), nil
}

func (d *imageDomain) ResizeImageToLimit(bin []byte) ([]byte, error) {
	// resize_to_limit
	// 縦横両方とも閾値より小さければそのままbinを返す
	input := bytes.NewReader(bin)
	var output bytes.Buffer
	cmd := exec.Command("convert", "-resize", strconv.Itoa(d.ResizeToLimit["height"]) + "x" + strconv.Itoa(d.ResizeToLimit["width"]), "-", "-")
	cmd.Stdin = input
	cmd.Stdout = &output
	if err := cmd.Run(); err != nil {
		return nil, xerrors.Errorf("Resize error: %w", err)
	}
	return output.Bytes(), nil
}

// http://noodles-mtb.hatenablog.com/entry/2013/07/08/151316
// 縦横比を保持したまま、指定されたサイズに収まるようリサイズします。
func (i *imageDomain) CompareImage(sourceBin []byte, targetBin []byte) (bool, error) {
	// magick_local_image = Magick::Image.from_blob(local_image_bin).first
	// magick_remote_image = Magick::Image.from_blob(remote_image_bin).first
	// local_small_image = magick_local_image.resize_to_fit(100)
	// remote_small_image = magick_remote_image.resize_to_fit(100)
	// diff = local_small_image.composite(remote_small_image, 0, 0, Magick::DifferenceCompositeOp)
	// diff.channel_mean.first.to_i <= 3500
	return true
}
