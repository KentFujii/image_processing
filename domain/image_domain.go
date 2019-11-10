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
	"fmt"
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
	var outputBrb bytes.Buffer
	cmd := exec.Command("convert", "-", d.ConvertTo + ":-")
	cmd.Stdin = bytes.NewReader(bin)
	cmd.Stdout = &outputBrb
	if err := cmd.Run(); err != nil {
		return nil, xerrors.Errorf("ConvertFormat error: %w", err)
	}
	return outputBrb.Bytes(), nil
}

func (d *imageDomain) ResizeImageToLimit(bin []byte) ([]byte, error) {
	var outputBrb bytes.Buffer
	size := strconv.Itoa(d.ResizeToLimit["height"]) + "x" + strconv.Itoa(d.ResizeToLimit["width"]) + ">"
	cmd := exec.Command("convert", "-", "-resize", size , "-")
	cmd.Stdin = bytes.NewReader(bin)
	cmd.Stdout = &outputBrb
	if err := cmd.Run(); err != nil {
		return nil, xerrors.Errorf("ResizeImageToLimit error: %w", err)
	}
	return outputBrb.Bytes(), nil
}

func (i *imageDomain) CompareImage(srcBin []byte, dstBin []byte) (bool, error) {
	_, srcFormat, _ := image.DecodeConfig(bytes.NewReader(srcBin))
	srcU, _ := uuid.NewRandom()
	inputSrcTempFile, _ := ioutil.TempFile(os.TempDir(), srcU.String() + "-compareImage-src-*" + "." + srcFormat)
	var outputSrcBrb bytes.Buffer
	srcCmd := exec.Command("convert", "-", "-resize", "100x100" , "-")
	srcCmd.Stdin = bytes.NewReader(srcBin)
	srcCmd.Stdout = &outputSrcBrb
	if srcCmdErr := srcCmd.Run(); srcCmdErr != nil {
		return false, xerrors.Errorf("CompareImage error: %w", srcCmdErr)
	}
	inputSrcTempFile.Write(outputSrcBrb.Bytes())
	defer os.Remove(inputSrcTempFile.Name())

	_, dstFormat, _ := image.DecodeConfig(bytes.NewReader(dstBin))
	dstU, _ := uuid.NewRandom()
	inputDstTempFile, _ := ioutil.TempFile(os.TempDir(), dstU.String() + "-compareImage-dst-*" + "." + dstFormat)
	var outputDstBrb bytes.Buffer
	dstCmd := exec.Command("convert", "-", "-resize", "100x100" , "-")
	dstCmd.Stdin = bytes.NewReader(dstBin)
	dstCmd.Stdout = &outputDstBrb
	if dstCmdErr := dstCmd.Run(); dstCmdErr != nil {
		fmt.Println(11111)
		return false, xerrors.Errorf("CompareImage error: %w", dstCmdErr)
	}
	inputDstTempFile.Write(outputDstBrb.Bytes())
	defer os.Remove(inputDstTempFile.Name())

	var resultBrb bytes.Buffer
	c1 := exec.Command("convert", "-compose", "difference", inputSrcTempFile.Name(), inputDstTempFile.Name(), "sample.jpg")
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
	fmt.Println(diff)

	diffF, _ := strconv.ParseFloat(diff, 64)
	diffBigF := big.NewFloat(diffF)
	thldBigF := big.NewFloat(3500.0)
	result := diffBigF.Cmp(thldBigF)
	if result < 0 {
		return true, nil
	}
	return false, nil
}
