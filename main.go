package main

import (
	"flag"
	"./converter"
)

var (
	srcDir = flag.String("src", ".", "source dir")
	from = flag.String("from", "jpg", "from image type")
	to = flag.String("to", "png", "to image type")
)
func main() {
	flag.Parse()

	target := converter.GetFiles(*srcDir)
	if target == nil {
		panic("Invalid src dir")
	}

	for _, file := range target.Files {
		err := converter.ConvertToPng(file)
		if err != nil {
			panic(err)
		}
		println(file)
	}
}
