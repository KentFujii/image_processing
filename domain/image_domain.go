package domain

import (
	"bytes"
	"os"
	"os/exec"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"math/big"
	"image"
	_ "image/jpeg"
	_ "image/gif"
	_ "image/png"
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
	inputBrb := bytes.NewReader(bin)
	_, format, _ := image.DecodeConfig(inputBrb)
	u, _ := uuid.NewRandom()
	inputTempFile, _ := ioutil.TempFile(os.TempDir(), u.String() + "-convertFormat-*" + "." + format)
	inputTempFile.Write(bin)
	defer os.Remove(inputTempFile.Name())
	var outputBrb bytes.Buffer
	cmd := exec.Command("convert", inputTempFile.Name(), d.ConvertTo + ":-")
	cmd.Stdout = &outputBrb
	if err := cmd.Run(); err != nil {
		return nil, xerrors.Errorf("ConvertFormat error: %w", err)
	}
	return outputBrb.Bytes(), nil
}

func (d *imageDomain) ResizeImageToLimit(bin []byte) ([]byte, error) {
	inputBrb := bytes.NewReader(bin)
	_, format, _ := image.DecodeConfig(inputBrb)
	u, _ := uuid.NewRandom()
	inputTempFile, _ := ioutil.TempFile(os.TempDir(), u.String() + "-resizeImageToLimit-*" + "." + format)
	inputTempFile.Write(bin)
	defer os.Remove(inputTempFile.Name())
	var outputBrb bytes.Buffer
	size := strconv.Itoa(d.ResizeToLimit["height"]) + "x" + strconv.Itoa(d.ResizeToLimit["width"]) + ">"
	cmd := exec.Command("convert", inputTempFile.Name(), "-resize", size , "-")
	cmd.Stdout = &outputBrb
	if err := cmd.Run(); err != nil {
		return nil, xerrors.Errorf("Resize error: %w", err)
	}
	return outputBrb.Bytes(), nil
}

func (i *imageDomain) CompareImage(srcBin []byte, dstBin []byte) (bool, error) {
	inputSrcBrb := bytes.NewReader(srcBin)
	_, srcFormat, _ := image.DecodeConfig(inputSrcBrb)
	srcU, _ := uuid.NewRandom()
	inputSrcTempFile, _ := ioutil.TempFile(os.TempDir(), srcU.String() + "-compareImage-src-*" + "." + srcFormat)
	inputSrcTempFile.Write(srcBin)
	defer os.Remove(inputSrcTempFile.Name())

	inputDstBrb := bytes.NewReader(dstBin)
	_, dstFormat, _ := image.DecodeConfig(inputDstBrb)
	dstU, _ := uuid.NewRandom()
	inputDstTempFile, _ := ioutil.TempFile(os.TempDir(), dstU.String() + "-compareImage-dst-*" + "." + dstFormat)
	inputDstTempFile.Write(dstBin)
	defer os.Remove(inputDstTempFile.Name())

	var resultBrb bytes.Buffer
	c1 := exec.Command("convert", "-compose", "difference", inputSrcTempFile.Name(), inputDstTempFile.Name(), "-")
	c2 := exec.Command("identify", "-format", "'%[mean]'", "-")
	pr, pw := io.Pipe()
	c1.Stdout = pw
	c2.Stdin = pr
	c2.Stdout = &resultBrb
	c1.Start()
	c2.Start()
	c1.Wait()
	pw.Close()
	c2.Wait()
	diff := strings.Trim(resultBrb.String(), `'`)

	diffF, _ := strconv.ParseFloat(diff, 64)
	diffBigF := big.NewFloat(diffF)
	thldBigF := big.NewFloat(3500.0)
	result := diffBigF.Cmp(thldBigF)
	if result < 0 {
		return true, nil
	}
	return false, nil
}
