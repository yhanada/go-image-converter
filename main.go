package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yhanada/go-image-converter/converter"
)

var (
	srcDir = flag.String("src", ".", "source dir")
	from   = flag.String("from", "jpg", "from image type")
	to     = flag.String("to", "png", "to image type")
)

func main() {
	flag.Parse()

	fromType := converter.GetImageType(*from)
	toType := converter.GetImageType(*to)
	if fromType == toType {
		fmt.Println("Error: Same image type:" + *from)
		os.Exit(1)
	}

	files := converter.GetTargetFiles(fromType, *srcDir)
	if len(files) == 0 {
		fmt.Println("No target files")
		os.Exit(0)
	}

	for _, file := range files {
		ok, dest, err := converter.Convert(toType, file)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if ok {
			fmt.Println("Converted:" + file + " -> " + dest)
		} else {
			fmt.Println("Failed to convert:" + file)
		}
	}
}
