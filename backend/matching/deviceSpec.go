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
		name:         "iPhone 6/7/8(s)",
		width:        1334,
		height:       750,
		startLeftX:   183,
		startLeftY:   206,
		startRightX:  693,
		startRightY:  206,
		recWidth:     72,
		recHeight:    128,
		recWidthStep: 93,
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
