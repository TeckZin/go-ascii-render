package renderer

import (
	"fmt"
	"main/internal/importer"
	"math"
)

type Color struct {
	R, G, B uint8
}

// convert at color to byte
type Pixel struct {
	red            int32
	blue           int32
	green          int32
	AsciiCharacter byte
}

type AsciiImage struct {
	PixelMap        [][]*Pixel
	ANSIEncodingMap [][]*string
}

func InitAsciiImage(imageData *importer.Image) (*AsciiImage, error) {

	asciiImage := &AsciiImage{}

	asciiImage.convertBrightnessToAscii(imageData.BrightnessMap, int32(imageData.Width), int32(imageData.Height))

	asciiImage.getImageColor(imageData)
	asciiImage.getAnsiEncoding()

	return asciiImage, nil

}

func rgbToANSI(rgb Color) string {
	r := uint8(math.Round(float64(rgb.R) / 51))
	g := uint8(math.Round(float64(rgb.G) / 51))
	b := uint8(math.Round(float64(rgb.B) / 51))
	return fmt.Sprintf("\033[38;5;%dm", 16+36*r+6*g+b)
}

func (a *AsciiImage) convertBrightnessToAscii(brightnessMap [][]int32, width int32, height int32) {
	// asciiCharacters := []string{"@", "$", "%", "#", "*", "+", "=", "-", ":", "."}

	asciiCharacters := []string{".", ":", "-", "=", "+", "*", "#", "%", "$", "@"}
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

// func (a *AsciiImage) convertGrayScaleToAscii(img *importer.Image) {
// 	// for _, g := range img.GrayScaleMap {
// 	// 	fmt.Println(g)
// 	//
// 	// }
// 	//
// 	asciiCharacters := []string{"@", "$", "%", "#", "*", "+", "=", "-", ":", "."}
// 	for y := 0; y < int(img.Height); y++ {
// 		var out string
// 		row := make([]*Pixel, 0)
// 		for x := 0; x < int(img.Width); x++ {
// 			ascii := asciiCharacters[int(img.GrayScaleMap[y][x])*(len(asciiCharacters)-1)/255]
// 			out = out + ascii
// 			pixel := &Pixel{}
// 			pixel.AsciiCharacter = ascii[0]
// 			row = append(row, pixel)
// 		}
// 		a.PixelMap = append(a.PixelMap, row)
// 		// fmt.Println(out)
//
// 	}
//
// }

func (a *AsciiImage) getImageColor(img *importer.Image) {
	for y, row := range img.RedColorMap {
		for x, r := range row {
			(*a.PixelMap[y][x]).red = r
			(*a.PixelMap[y][x]).blue = img.BlueColorMap[y][x]
			(*a.PixelMap[y][x]).green = img.GreenColorMap[y][x]
		}
	}

}

func (a *AsciiImage) getAnsiEncoding() {

	ansiMap := make([][]*string, 0)
	for _, row := range a.PixelMap {
		ansiRow := make([]*string, 0)
		for _, pixel := range row {
			rgb := Color{R: uint8(pixel.red), G: uint8(pixel.green), B: uint8(pixel.blue)}
			ansi := rgbToANSI(rgb)
			ansiRow = append(ansiRow, &ansi)
		}
		ansiMap = append(ansiMap, ansiRow)

	}

	a.ANSIEncodingMap = ansiMap

}

func (a *AsciiImage) Display() {
	for y, pixelRow := range a.PixelMap {
		for x, pixel := range pixelRow {
			fmt.Print(string(*a.ANSIEncodingMap[y][x]))
			fmt.Print(string(pixel.AsciiCharacter))
			fmt.Print("\033[0m")

		}
		fmt.Println()
	}

}
