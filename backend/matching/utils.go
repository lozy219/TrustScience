package matching

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"strings"

	"golang.org/x/image/draw"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
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

func LoadImage(file multipart.File) image.Image {
	// hack hack hack
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	str := buf.String()
	file.Seek(0, 0)

	var img image.Image
	var err error

	if strings.HasPrefix(str, "\xff\xd8\xff") {
		img, err = jpeg.Decode(file)
		checkErr(err)
	} else if strings.HasPrefix(str, "\x89PNG\r\n\x1a\n") {
		img, err = png.Decode(file)
		checkErr(err)
	} else if strings.HasPrefix(str, "GIF8") {
		img, err = gif.Decode(file)
		checkErr(err)
	} else {
		fmt.Println("invalid image type")
	}

	return img
}
