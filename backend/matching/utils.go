package matching

import (
	"crypto/md5"
	"fmt"
	"image"
	"image/png"

	"golang.org/x/image/draw"
)

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func crop(src image.Image, rec image.Rectangle) image.Image {
	result := image.NewGray(rec)
	draw.Draw(result, rec, src, rec.Min, draw.Over)

	return result
}

func convertToGray(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Max.X, bounds.Max.Y
	result := image.NewGray(image.Rect(0, 0, w, h))
	draw.Draw(result, result.Bounds(), src, bounds.Min, draw.Over)

	return result
}

func scaleImage(src image.Image, targetSize image.Point) image.Image {
	result := image.NewGray(image.Rect(0, 0, targetSize.X, targetSize.Y))
	draw.NearestNeighbor.Scale(result, result.Bounds(), src, src.Bounds(), draw.Src, nil)

	return result
}

func HashImage(src image.Image) string {
	h := md5.New()
	err := png.Encode(h, src)
	checkErr(err)

	return fmt.Sprintf("%x", h.Sum(nil))
}
