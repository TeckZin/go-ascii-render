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

	imageFile.ScaleImageRation(0.0125)
	// imageFile.GetBrightness()
	imageFile.GetGrayScale()

	// fmt.Println(*imageFile.ImageValue)
	// fmt.Println(imageFile.BrightnessMap)
	// fmt.Println((*imageFile.ImageValue).ColorModel())

	asciiImage, err := renderer.InitAsciiImage(imageFile)

	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(asciiImage)

	for _, x := range asciiImage.PixelMap {
		for _, y := range x {
			fmt.Print(string(y.AsciiCharacter))
		}
		fmt.Println()
	}

	// for _, x := range imageFile.BrightnessMap {
	// 	for _, y := range x {
	// 		if len(string(y)) != 2 {
	//
	// 			fmt.Print(y, "  ")
	// 		} else {
	// 			fmt.Print(y, " ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }

}
