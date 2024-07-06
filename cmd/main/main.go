package main

import (
	"fmt"
	"main/internal/importer"
	"main/internal/renderer"
)

func main() {
	path := "internal/images/test1.png"
	fmt.Println(path)

	imageFile := &importer.Image{}

	imageFile.Path = path

	if err := imageFile.ReadImage(); err != nil {
		fmt.Println(err)
		return
	}

	imageFile.ScaleImageRation(0.18)
	imageFile.GetColorMap()
	// imageFile.GetGrayScale()
	imageFile.GetBrightness()

	asciiImage, err := renderer.InitAsciiImage(imageFile)

	if err != nil {
		fmt.Println(err)
		return
	}

	asciiImage.Display()

}
