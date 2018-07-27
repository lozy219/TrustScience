package matching

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"sync"
)

func calcDiffV3(x uint32, y uint32) uint32 {
	if x < y {
		return y - x
	} else {
		return x - y
	}
}

func crop(src image.Image, rec image.Rectangle) image.Image {
	result := image.NewGray(rec)
	draw.Draw(result, rec, src, rec.Min, draw.Over)

	return result
}

func findDiff(src image.Image, target image.Image) uint32 {
	bounds := src.Bounds()

	tBounds := target.Bounds()
	w, h := tBounds.Max.X, tBounds.Max.Y

	var sumDiff uint32

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			color := src.At(bounds.Min.X+x, bounds.Min.Y+y)
			g, _, _, _ := color.RGBA()

			tColor := target.At(x, y)
			tg, _, _, _ := tColor.RGBA()
			diff := calcDiffV3(g, tg)
			sumDiff += diff
		}
	}

	return sumDiff
}

func findMatch(src image.Image) string {
	targetFolder := "avatars/"
	dir, err := os.Open(targetFolder)
	checkErr(err)
	defer dir.Close()

	var minDiff uint32 = 111111111
	minDiffName := ""

	var wg sync.WaitGroup
	list, _ := dir.Readdirnames(0)
	for _, name := range list {
		wg.Add(1)
		go func(name string) {
			cursor := len(name) - 4
			if len(name) > 4 && name[cursor:] == ".PNG" {
				target := targetFolder + name
				tInfile, err := os.Open(target)
				checkErr(err)
				tSrc, err := png.Decode(tInfile)
				checkErr(err)

				diff := findDiff(src, tSrc)
				if diff < minDiff {
					minDiff = diff
					minDiffName = name
				}
			}
			wg.Done()
		}(name)
	}
	wg.Wait()

	return minDiffName
}

func MatchV3(path string) []string {
	infile, err := os.Open(path)
	checkErr(err)
	defer infile.Close()

	src, err := png.Decode(infile)
	checkErr(err)

	src = convertToGray(src)

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

		fmt.Println(findMatch(crop(src, recLeft)))
		fmt.Println(findMatch(crop(src, recRight)))
	}

	lst := make([]string, 10)

	return lst
}
