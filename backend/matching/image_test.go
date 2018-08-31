package matching

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func loadImage() image.Image {
	target := "test/screenshots/test.PNG"
	tInfile, err := os.Open(target)
	checkErr(err)
	tSrc, err := png.Decode(tInfile)
	checkErr(err)

	return tSrc
}

func TestImageScale(t *testing.T) {
	targetSize := image.Point{160, 160}
	src := loadImage()
	result := scaleImage(src, targetSize)

	fname := "./test/out/testScaleOut.png"
	fout, err := os.Create(fname)
	checkErr(err)
	defer fout.Close()

	encodeErr := png.Encode(fout, result)
	checkErr(encodeErr)

	if result.Bounds().Size() != targetSize {
		t.Errorf("Scale size test failed.")
	}
}

func TestImageConvertToGray(t *testing.T) {

}

func TestImageHash(t *testing.T) {
	src := loadImage()
	hash := HashImage(src)
	expected := "ef095c76e6581c9ebc663b327ef98351"

	if hash != expected {
		t.Errorf("hash %s is different from expected %s", hash, expected)
	}
}
