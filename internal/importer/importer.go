package importer

import (
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"
)

type Image struct {
	Path          string
	Width         int
	Height        int
	BrightnessMap [][]int32
	GrayScaleMap  [][]int32
	RedColorMap   [][]int32
	BlueColorMap  [][]int32
	GreenColorMap [][]int32
	ImageValue    *image.Image
}

func (img *Image) ReadImage() error {
	file, err := os.Open(img.Path)
	if err != nil {
		return err
	}

	defer file.Close()

	imageDec, _, err := image.Decode(file)

	img.ImageValue = &imageDec

	bounds := imageDec.Bounds()

	img.Height = bounds.Max.Y
	img.Width = bounds.Max.X

	if err != nil || img.ImageValue != nil {
		return err
	}

	return nil

}

func (img *Image) ScaleImageRation(ratio float32) {
	newWidth := int32(float32(img.Width) * ratio / 2)
	newHeight := int32(float32(img.Height) * ratio / 3)

	img.ScaleImageBounds(newWidth, newHeight)

}
func (img *Image) ScaleImageBounds(width int32, height int32) {
	img.Width = int(width)
	img.Height = int(height)
	emptyImage := image.NewRGBA(image.Rect(0, 0, int(width), int(height)))
	draw.ApproxBiLinear.Scale(emptyImage, emptyImage.Bounds(), *img.ImageValue, (*img.ImageValue).Bounds(), draw.Over, nil)

	newImg := image.Image(emptyImage)
	img.ImageValue = &newImg

}

func (img *Image) GetColorMap() {
	redMap := make([][]int32, 0)
	greenMap := make([][]int32, 0)
	blueMap := make([][]int32, 0)
	for y := 0; y < img.Height; y++ {
		redRow := make([]int32, 0)
		greenRow := make([]int32, 0)
		blueRow := make([]int32, 0)
		for x := 0; x < img.Width; x++ {
			cords := (*img.ImageValue).At(x, y)
			r, g, b, _ := cords.RGBA()
			redRow = append(redRow, int32(r))
			greenRow = append(greenRow, int32(g))
			blueRow = append(blueRow, int32(b))

		}
		redMap = append(redMap, redRow)
		greenMap = append(greenMap, greenRow)
		blueMap = append(blueMap, blueRow)
	}

	img.RedColorMap = redMap
	img.GreenColorMap = greenMap
	img.BlueColorMap = blueMap

}

// func (img *Image) GetBrightness() {
// 	newBrightnessMap := make([][]int32, 0)
// 	for y := 0; y < img.Height; y++ {
// 		row := make([]int32, 0)
// 		for x := 0; x < img.Width; x++ {
// 			r, g, b, _ := (*img.ImageValue).At(x, y).RGBA()
// 			brightness := (0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)) / 257
// 			row = append(row, int32(brightness))
// 		}
// 		newBrightnessMap = append(newBrightnessMap, row)
// 	}
// 	img.BrightnessMap = newBrightnessMap
//
// }

func (img *Image) GetGrayScale() {
	newGrayScaleMap := make([][]int32, 0)
	for y := 0; y < img.Height; y++ {
		grayScaleRow := make([]int32, 0)
		for x := 0; x < img.Width; x++ {
			c := color.GrayModel.Convert((*img.ImageValue).At(x, y))
			grayScaleRow = append(grayScaleRow, int32(c.(color.Gray).Y))
		}
		newGrayScaleMap = append(newGrayScaleMap, grayScaleRow)
	}
	img.GrayScaleMap = newGrayScaleMap

}
