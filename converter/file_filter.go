package converter

import (
	"io/ioutil"
	"path/filepath"
)

type Target struct {
	Files []string
}

func GetFiles(srcDir string) *Target {
	files := dirWalk(srcDir)
	return &Target{Files: files}
}

func dirWalk(dir string) [] string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, dirWalk(filepath.Join(dir, file.Name()))...)
		} else {
			paths = append(paths, filepath.Join(dir, file.Name()))
		}
	}

	return paths
}