package image_lib

import (
	"image"
)

type Matrix interface {
	Matrix(src image.Image, width int, height int)
	Image() image.Image
	Width() int
	Height() int
}
