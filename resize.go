package resizeimage

import (
	"image"

	"github.com/liontail/imaging"
)

func Resize(fileName, destinationFile string, height, width int) (*image.NRGBA, error) {
	tempOpen, err := imaging.Open(fileName)
	if err != nil {
		return nil, err
	}

	return imaging.Resize(tempOpen, height, width, imaging.Lanczos), nil
}

func SaveImageToDestination(dstImage *image.NRGBA, destinationFile string) error {
	return imaging.Save(dstImage, destinationFile)
}
