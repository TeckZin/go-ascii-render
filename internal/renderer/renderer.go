package renderer

import (
	"fmt"
	"main/internal/importer"
)

type Pixel struct {
	red            int32
	blue           int32
	green          int32
	AsciiCharacter byte
}

type AsciiImage struct {
	PixelMap [][]*Pixel
}

func InitAsciiImage(imageData *importer.Image) (*AsciiImage, error) {
	// fmt.Println(imageData.RedColorMap)

	asciiImage := &AsciiImage{}

	// asciiImage.convertBrightnessToAscii(imageData.BrightnessMap, int32(imageData.Width), int32(imageData.Height))

	asciiImage.converGrayScaleToAscii(imageData)

	return asciiImage, nil

}

func (a *AsciiImage) convertBrightnessToAscii(brightnessMap [][]int32, width int32, height int32) {
	for _, b := range brightnessMap {
		fmt.Println(b)
	}

	asciiCharacters := []string{"@", "$", "%", "#", "*", "+", "=", "-", ":", "."}

	// asciiCharacters := []string{".", ":", "-", "=", "+", "*", "#", "%", "$", "@"}
	for y := 0; y < int(height); y++ {
		var out string
		row := make([]*Pixel, 0)
		for x := 0; x < int(width); x++ {
			ascii := asciiCharacters[int(brightnessMap[y][x])*(len(asciiCharacters)-1)/255]
			out = out + ascii
			pixel := &Pixel{}
			pixel.AsciiCharacter = ascii[0]
			row = append(row, pixel)
		}
		a.PixelMap = append(a.PixelMap, row)
		// fmt.Println(out)

	}

}

func (a *AsciiImage) converGrayScaleToAscii(img *importer.Image) {
	for _, g := range img.GrayScaleMap {
		fmt.Println(g)

	}

	asciiCharacters := []string{"@", "$", "%", "#", "*", "+", "=", "-", ":", "."}
	for y := 0; y < int(img.Height); y++ {
		var out string
		row := make([]*Pixel, 0)
		for x := 0; x < int(img.Width); x++ {
			ascii := asciiCharacters[int(img.GrayScaleMap[y][x])*(len(asciiCharacters)-1)/255]
			out = out + ascii
			pixel := &Pixel{}
			pixel.AsciiCharacter = ascii[0]
			row = append(row, pixel)
		}
		a.PixelMap = append(a.PixelMap, row)
		// fmt.Println(out)

	}

}
