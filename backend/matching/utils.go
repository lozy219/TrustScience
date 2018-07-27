package matching

import (
	"image"
	"image/draw"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func convertToGray(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	result := image.NewGray(image.Rect(0, 0, w, h))
	draw.Draw(result, result.Bounds(), src, bounds.Min, draw.Over)

	return result
}
