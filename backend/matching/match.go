package matching

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/lozy219/trustscience/backend/record"
)

var cachedImages = map[string]image.Image{}

func calcDiff(x uint32, y uint32) uint32 {
	if x < y {
		return y - x
	} else {
		return x - y
	}
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
			diff := calcDiff(g, tg)
			sumDiff += diff
		}
	}

	return sumDiff
}

func findMatch(src image.Image) string {
	var minDiff uint32 = 100000000
	minDiffName := ""

	for name, tSrc := range cachedImages {
		diff := findDiff(src, tSrc)
		if diff < minDiff {
			minDiff = diff
			minDiffName = name
		}
	}
	return minDiffName
}

func Match(src image.Image) []string {
	src = convertToGray(src)
	lst := make([]string, 10)

	spec := GetDeviceSpec(src.Bounds().Size())
	if spec == nil {
		src = scaleImage(src, image.Point{1334, 750})
	} else {
		// try to force scale it down to iPhone 6 size
		if spec.shouldCrop {
			src = crop(src, image.Rect(spec.cropLeftX, spec.cropLeftY, spec.cropRightX, spec.cropRightY))
		}
		if spec.shouldResize {
			src = scaleImage(src, image.Point{1334, 750})
		}
	}
	// fname := "./matching/test/out/testMatch.png"
	// fout, err := os.Create(fname)
	// checkErr(err)
	// defer fout.Close()
	// encodeErr := png.Encode(fout, src)
	// checkErr(encodeErr)

	spec = GetDefaultDeviceSpec()
	startLeft := spec.StartLeft()
	startRight := spec.StartRight()
	widthStep := spec.recWidthStep
	matchRect := spec.MatchRect()

	isResultValid := true

	for i := 0; i < 5; i++ {
		x0Left := startLeft.X + i*widthStep
		y0Left := startLeft.Y
		x1Left := x0Left + matchRect.X
		y1Left := y0Left + matchRect.Y
		recLeft := image.Rect(x0Left, y0Left, x1Left, y1Left)

		x0Right := startRight.X + i*widthStep
		y0Right := startRight.Y
		x1Right := x0Right + matchRect.X
		y1Right := y0Right + matchRect.Y
		recRight := image.Rect(x0Right, y0Right, x1Right, y1Right)

		matchLeft := findMatch(crop(src, recLeft))
		matchRight := findMatch(crop(src, recRight))

		if len(matchLeft) == 0 {
			lst[i] = "不知道"
			isResultValid = false
		} else {
			lst[i] = matchLeft[:len(matchLeft)-4]
		}

		if len(matchRight) == 0 {
			lst[i+5] = "不知道"
			isResultValid = false
		} else {
			lst[i+5] = matchRight[:len(matchRight)-4]
		}
	}

	if isResultValid {
		record.NewRecord(lst)
	}

	fmt.Print(lst)

	return lst
}

func init() {
	targetFolder := "avatars/"
	dir, err := os.Open(targetFolder)
	checkErr(err)
	defer dir.Close()

	list, _ := dir.Readdirnames(0)
	for _, name := range list {
		cursor := len(name) - 4
		if len(name) > 4 && name[cursor:] == ".PNG" {
			target := targetFolder + name
			tInfile, err := os.Open(target)
			checkErr(err)
			defer tInfile.Close()
			tSrc, err := png.Decode(tInfile)
			checkErr(err)
			cachedImages[name] = tSrc
		}
	}
}
