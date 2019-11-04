package domain

import (
	"bytes"
	"os/exec"
	"strconv"
	"fmt"
	"golang.org/x/xerrors"
	"github.com/google/uuid"
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
	input := bytes.NewReader(bin)
	var output bytes.Buffer
	var stderr bytes.Buffer
	// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/functions/imagemagick/imagemagick.go#L93
	size := strconv.Itoa(d.ResizeToLimit["height"]) + "x" + strconv.Itoa(d.ResizeToLimit["width"])
	cmd := exec.Command("convert", "-resize", size, "-", "-")
	cmd.Stdin = input
	cmd.Stdout = &output
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		fmt.Println(11111)
		// fmt.Println(cmd.Stdin)
		fmt.Println(cmd.Args)
		fmt.Println(stderr.String())
		fmt.Println(err)
		return nil, xerrors.Errorf("Resize error: %w", err)
	}
	return output.Bytes(), nil
}

// https://hawksnowlog.blogspot.com/2019/04/generate-uuid-with-golang.html
// https://golang.org/pkg/io/ioutil/#TempFile
// 安全のためstdinは全部Tempfileを経由する、Tempfileの場所をConfigに追加
// root@505fec980135:/go/src# cat domain/testdata/jpeg/butterfly-100kb.jpg | convert -resize 600x600 - - | identify -
// -=>/tmp/magick-39630KRPu1RZEhUEX JPEG 600x600 600x600+0+0 8-bit sRGB 79069B 0.000u 0:00.000
// http://noodles-mtb.hatenablog.com/entry/2013/07/08/151316
// 縦横比を保持したまま、指定されたサイズに収まるようリサイズします。
// https://qiita.com/kwst/items/c40817b3cdf841995257
// https://rmagick.github.io/image1.html#composite
// dest.composite(src, x, y, composite_op) -> image
// composite -compose difference domain/testdata/jpeg/butterfly-100kb.jpg domain/testdata/jpeg/butterfly-500kb.jpg sample.jpeg
// cat domain/testdata/jpeg/butterfly-100kb.jpg | composite -compose difference - domain/testdata/jpeg/butterfly-500kb.jpg sample.jpg
func (i *imageDomain) CompareImage(srcBin []byte, dstBin []byte) (bool, error) {
	// magick_local_image = Magick::Image.from_blob(local_image_bin).first
	// magick_remote_image = Magick::Image.from_blob(remote_image_bin).first
	// local_small_image = magick_local_image.resize_to_fit(100)
	// remote_small_image = magick_remote_image.resize_to_fit(100)
	// diff = local_small_image.composite(remote_small_image, 0, 0, Magick::DifferenceCompositeOp)
	// diff.channel_mean.first.to_i <= 3500
	u, _ := uuid.NewRandom()
	fmt.Println(u)
	fmt.Println(srcBin)
	fmt.Println(dstBin)
	// srcTmpfile, err := ioutil.TempFile("/tmp", "example")
	// sourceTmpfile, err := ioutil.TempFile("", "example")
	// input := bytes.NewReader(bin)
	return true, nil
}
