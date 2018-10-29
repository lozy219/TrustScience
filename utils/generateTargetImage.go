package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/corona10/goimagehash"
	"golang.org/x/image/draw"
)

func main() {
	targetFolder := "../data/matches/"
	outputFolder := "../data/avatars/"

	dir, err := os.Open(targetFolder)
	checkErr(err)
	defer dir.Close()

	list, _ := dir.Readdirnames(0)
	for _, name := range list {
		cursor := len(name) - 4
		if cursor > 0 && name[cursor:] == ".PNG" {
			matchPath := targetFolder + name
			matchFile, err := os.Open(matchPath)
			checkErr(err)
			matchSrc, err := png.Decode(matchFile)
			checkErr(err)

			resizedSrc := image.NewGray(image.Rect(0, 0, 1334, 750))
			draw.NearestNeighbor.Scale(resizedSrc, resizedSrc.Bounds(), matchSrc, matchSrc.Bounds(), draw.Src, nil)

			generateOutput(resizedSrc, name, outputFolder)
		}
	}

}

func generateOutput(src image.Image, srcName string, outputFolder string) {
	startLeft := image.Point{183, 206}
	startRight := image.Point{693, 206}
	widthStep := 93
	recWidth := 72
	recHeight := 128

	for i := 0; i < 5; i++ {
		x0Left := startLeft.X + i*widthStep
		y0Left := startLeft.Y
		x1Left := x0Left + recWidth
		y1Left := y0Left + recHeight
		recLeft := image.Rect(x0Left, y0Left, x1Left, y1Left)

		x0Right := startRight.X + i*widthStep
		y0Right := startRight.Y
		x1Right := x0Right + recWidth
		y1Right := y0Right + recHeight
		recRight := image.Rect(x0Right, y0Right, x1Right, y1Right)

		hashAndSave(crop(src, recLeft), srcName, outputFolder)
		hashAndSave(crop(src, recRight), srcName, outputFolder)
	}
}

func crop(src image.Image, rec image.Rectangle) image.Image {
	result := image.NewGray(rec)
	draw.Draw(result, rec, src, rec.Min, draw.Over)

	return result
}

func hashAndSave(src image.Image, srcName string, outputFolder string) {
	hash, err := goimagehash.AverageHash(src)
	checkErr(err)
	destination := fmt.Sprintf("%s%s_%s", outputFolder, fmt.Sprintf("%v", hash.GetHash()), srcName)
	out, err := os.Create(destination)
	checkErr(err)
	defer out.Close()

	err = png.Encode(out, src)
	checkErr(err)
}

func checkErr(e error) {
	if e != nil {
		panic(e.Error())
	}
}
