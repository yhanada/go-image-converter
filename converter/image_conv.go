package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
)

// Convert image from src file to image specified with targetType
func Convert(targetType ImageType, src string) (bool, string, error) {
	file, err := os.Open(src)
	if err != nil {
		return false, "", err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return false, "", err
	}

	destFilepath, err := getDestFilepath(targetType, src)
	if err != nil {
		return false, "", err
	}
	dest, err := os.Create(destFilepath)
	defer dest.Close()

	switch targetType {
	case JPEG:
		jpeg.Encode(dest, img, nil)
	case PNG:
		png.Encode(dest, img)
	case GIF:
		gif.Encode(dest, img, nil)
	}
	return true, destFilepath, nil
}
