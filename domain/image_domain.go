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

// https://github.com/KentFujii/image_processing/commit/c8bc63ae30b75e69731d1879f343668b0a1e7119#diff-bb04ee0160ff1fefc1c86397621d9d45
// 全てtmpでやり取りする
func (i *imageDomain) CompareImage(srcBin []byte, dstBin []byte) (bool, error) {
	_, srcFormat, _ := image.DecodeConfig(bytes.NewReader(srcBin))
	inputSrcU, _ := uuid.NewRandom()
	inputSrcTempFile, _ := ioutil.TempFile(os.TempDir(), inputSrcU.String() + "-compareImage-inputSrc-*" + "." + srcFormat)
	inputSrcTempFile.Write(srcBin)
	var outputSrcBrb bytes.Buffer
	srcCmd := exec.Command("convert", inputSrcTempFile.Name(), "-resize", "100x100" , "-")
	srcCmd.Stdout = &outputSrcBrb
	if srcCmdErr := srcCmd.Run(); srcCmdErr != nil {
		return false, xerrors.Errorf("CompareImage error: %w", srcCmdErr)
	}
	outputSrcU, _ := uuid.NewRandom()
	outputSrcTempFile, _ := ioutil.TempFile(os.TempDir(), outputSrcU.String() + "-compareImage-outputSrc-*" + "." + srcFormat)
	outputSrcTempFile.Write(outputSrcBrb.Bytes())
	defer os.Remove(inputSrcTempFile.Name())
	defer os.Remove(outputSrcTempFile.Name())

	_, dstFormat, _ := image.DecodeConfig(bytes.NewReader(dstBin))
	inputDstU, _ := uuid.NewRandom()
	inputDstTempFile, _ := ioutil.TempFile(os.TempDir(), inputDstU.String() + "-compareImage-inputDst-*" + "." + dstFormat)
	inputDstTempFile.Write(dstBin)
	var outputDstBrb bytes.Buffer
	dstCmd := exec.Command("convert", inputDstTempFile.Name(), "-resize", "100x100" , "-")
	dstCmd.Stdout = &outputDstBrb
	if dstCmdErr := dstCmd.Run(); dstCmdErr != nil {
		return false, xerrors.Errorf("CompareImage error: %w", dstCmdErr)
	}
	outputDstU, _ := uuid.NewRandom()
	outputDstTempFile, _ := ioutil.TempFile(os.TempDir(), outputDstU.String() + "-compareImage-outputDst-*" + "." + dstFormat)
	outputDstTempFile.Write(outputDstBrb.Bytes())
	defer os.Remove(inputDstTempFile.Name())
	defer os.Remove(outputDstTempFile.Name())

	var resultBrb bytes.Buffer
	compositeCmd := exec.Command("composite", "-compose", "difference", outputSrcTempFile.Name(), outputDstTempFile.Name(), "-")
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
