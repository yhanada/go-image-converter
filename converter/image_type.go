package converter

import (
	"errors"
	"path/filepath"
	"strings"
)

type ImageType int

const (
	NONE = iota
	JPEG
	PNG
	GIF
)

// get ImageType by string
func GetImageType(str string) ImageType {
	value := strings.ToLower(str)
	switch value {
	case "jpg", "jpeg":
		return JPEG
	case "png":
		return PNG
	case "gif":
		return GIF
	default:
		return NONE
	}
}

// get ImageType by file extention
func getFileImageType(filename string) ImageType {
	ext := filepath.Ext(filename)
	str := ext[1:]
	return GetImageType(str)
}

// get destination file path
func getDestFilepath(targetType ImageType, path string) (string, error) {
	ext := filepath.Ext(path)
	var targetExt string
	switch targetType {
	case JPEG:
		targetExt = ".jpg"
	case PNG:
		targetExt = ".png"
	case GIF:
		targetExt = ".gif"
	default:
		return "", errors.New("Invalid Image type")
	}

	return strings.Replace(path, ext, targetExt, 1), nil
}
