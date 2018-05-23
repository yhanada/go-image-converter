package converter

import (
	"os"
	"image"
	"image/png"
	_ "image/jpeg"
	"path/filepath"
	"strings"
)

func ConvertToPng(src string) (err error) {
	file, err := os.Open(src)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil
	}

	ext := filepath.Ext(src)
	destFilename := strings.Replace(src, ext, ".png", 1)
	dest, err := os.Create(destFilename)
	defer dest.Close()

	png.Encode(dest, img)
	return nil
}
