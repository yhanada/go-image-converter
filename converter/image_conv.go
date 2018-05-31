package converter

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// Convert image from src file to image specified with targetType
func Convert(targetType ImageType, src string) (string, error) {
	file, err := os.Open(src)
	if err != nil {
		return "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return "", err
	}

	destFilepath, err := getDestFilepath(targetType, src)
	if err != nil {
		return "", err
	}
	dest, err := os.Create(destFilepath)
	if err != nil {
		return "", err
	}
	defer dest.Close()

	switch targetType {
	case JPEG:
		jpeg.Encode(dest, img, nil)
	case PNG:
		png.Encode(dest, img)
	case GIF:
		gif.Encode(dest, img, nil)
	default:
		return "", errors.New("Invalid target ImageType")
	}
	return destFilepath, nil
}
