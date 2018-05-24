package converter

import (
	"io/ioutil"
	"path/filepath"
)

// get target file list
func GetTargetFiles(imageType ImageType, srcDir string) []string {
	return findFiles(imageType, srcDir)
}

func findFiles(imageType ImageType, dir string) []string {
	var paths []string
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return paths
	}

	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, findFiles(imageType, filepath.Join(dir, file.Name()))...)
		} else {
			fileImageType := getFileImageType(file.Name())
			if imageType == fileImageType {
				paths = append(paths, filepath.Join(dir, file.Name()))
			}
		}
	}
	return paths
}
