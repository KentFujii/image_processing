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
		return nil, xerrors.Errorf("Resize error: %w", err)
	}
	return outputBrb.Bytes(), nil
}

func (i *imageDomain) CompareImage(srcBin []byte, dstBin []byte) (bool, error) {
	var srcBrb bytes.Buffer
	srcCmd := exec.Command("convert", "-", "-resize", "100x100" , "-")
	srcCmd.Stdin = bytes.NewReader(srcBin)
	srcCmd.Stdout = &srcBrb
	if srcCmdErr := srcCmd.Run(); srcCmdErr != nil {
		return false, xerrors.Errorf("CompareImage error: %w", srcCmdErr)
	}
	_, srcFormat, _ := image.DecodeConfig(bytes.NewReader(srcBin))
	srcU, _ := uuid.NewRandom()
	srcTempFile, _ := ioutil.TempFile(os.TempDir(), srcU.String() + "-compareImage-outputSrc-*" + "." + srcFormat)
	srcTempFile.Write(srcBrb.Bytes())
	defer os.Remove(srcTempFile.Name())

	var dstBrb bytes.Buffer
	dstCmd := exec.Command("convert", "-", "-resize", "100x100" , "-")
	dstCmd.Stdin = bytes.NewReader(dstBin)
	dstCmd.Stdout = &dstBrb
	if dstCmdErr := dstCmd.Run(); dstCmdErr != nil {
		return false, xerrors.Errorf("CompareImage error: %w", dstCmdErr)
	}
	_, dstFormat, _ := image.DecodeConfig(bytes.NewReader(dstBin))
	dstU, _ := uuid.NewRandom()
	dstTempFile, _ := ioutil.TempFile(os.TempDir(), dstU.String() + "-compareImage-outputDst-*" + "." + dstFormat)
	dstTempFile.Write(dstBrb.Bytes())
	defer os.Remove(dstTempFile.Name())

	var resultBrb bytes.Buffer
	compositeCmd := exec.Command("composite", "-compose", "difference", srcTempFile.Name(), dstTempFile.Name(), "-")
	identifyCmd := exec.Command("identify", "-format", "'%[mean]'", "-")
	pr, pw := io.Pipe()
	compositeCmd.Stdout = pw
	identifyCmd.Stdin = pr
	identifyCmd.Stdout = &resultBrb
	compositeCmd.Start()
	identifyCmd.Start()
	compositeCmd.Wait()
	pw.Close()
	identifyCmd.Wait()
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
