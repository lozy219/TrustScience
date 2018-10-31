package matching

import "image"

type DeviceSpec struct {
	name         string
	width        int
	height       int
	startLeftX   int
	startLeftY   int
	startRightX  int
	startRightY  int
	recWidth     int
	recHeight    int
	recWidthStep int
	cropLeftX    int
	cropLeftY    int
	cropRightX   int
	cropRightY   int
	isDefault    bool
	shouldCrop   bool
	shouldResize bool
}

func (d *DeviceSpec) Size() image.Point {
	return image.Point{d.width, d.height}
}

func (d *DeviceSpec) StartLeft() image.Point {
	return image.Point{d.startLeftX, d.startLeftY}
}

func (d *DeviceSpec) StartRight() image.Point {
	return image.Point{d.startRightX, d.startRightY}
}

func (d *DeviceSpec) MatchRect() image.Point {
	return image.Point{d.recWidth, d.recHeight}
}

var specs = []DeviceSpec{
	DeviceSpec{
		name:         "iPhone 6",
		width:        1334,
		height:       750,
		startLeftX:   183,
		startLeftY:   206,
		startRightX:  693,
		startRightY:  206,
		recWidth:     72,
		recHeight:    128,
		recWidthStep: 93,
		shouldCrop:   false,
		shouldResize: false,
	},
	DeviceSpec{
		name:         "QHD",
		width:        2560,
		height:       1440,
		isDefault:    false,
		shouldCrop:   false,
		shouldResize: true,
	},
	DeviceSpec{
		name:         "iPhone 6 Plus",
		width:        1920,
		height:       1080,
		isDefault:    false,
		shouldCrop:   false,
		shouldResize: true,
	},
	DeviceSpec{
		name:         "2208",
		width:        2208,
		height:       1242,
		isDefault:    false,
		shouldCrop:   false,
		shouldResize: true,
	},
	DeviceSpec{
		name:         "iPhone X",
		width:        2436,
		height:       1125,
		isDefault:    false,
		shouldCrop:   true,
		cropLeftX:    218,
		cropLeftY:    0,
		cropRightX:   2218,
		cropRightY:   1125,
		shouldResize: true,
	},
	DeviceSpec{
		name:         "iPhone XR",
		width:        1792,
		height:       828,
		isDefault:    false,
		shouldCrop:   true,
		cropLeftX:    160,
		cropLeftY:    0,
		cropRightX:   1632,
		cropRightY:   828,
		shouldResize: true,
	},
}

func GetDeviceSpec(size image.Point) *DeviceSpec {
	for _, spec := range specs {
		if spec.Size() == size {
			return &spec
		}
	}

	return nil
}

func GetDefaultDeviceSpec() *DeviceSpec {
	return &specs[0]
}
