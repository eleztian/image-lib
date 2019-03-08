package image_lib

import (
	"image"
	"image/color"
)

type GaryMatrix struct {
	matrix [][]uint8
}

// create a new rgba matrix from Image.
func (rm *GaryMatrix) Matrix(src image.Image, width int, height int) {
	//rm.matrix = make([][]uint8, height)
	//dst := convertToNRGBA(src)
	//for i := 0; i < height; i++ {
	//	rm.matrix[i] = make([]uint8, width)
	//	for j := 0; j < width; j++ {
	//		//r, g, b, a := dst.At(j, i).RGBA()
	//		rm.matrix[i][j] = 1
	//	}
	//}
	return
}

func (rm *GaryMatrix) Image() image.Image {
	height := rm.Height()
	width := rm.Width()
	garyImg := image.NewGray(image.Rect(0, 0, width, height))
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			garyImg.SetGray(j, i, color.Gray{Y: rm.matrix[i][j]})
		}
	}
	return garyImg
}

func (rm *GaryMatrix) Width() int {
	if rm.matrix == nil && len(rm.matrix) == 0 {
		return 0
	}
	return len(rm.matrix[0])
}

func (rm *GaryMatrix) Height() int {
	if rm.matrix == nil {
		return 0
	}
	return len(rm.matrix)
}
