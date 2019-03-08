package image_lib

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

// save a image matrix as a jpeg,if unsuccessful it will return a error,quality must be 1 to 100.
func saveAsJPEG(writer io.Writer, matrix Matrix, quality int) error {
	if quality < 1 {
		quality = 1
	} else if quality > 100 {
		quality = 100
	}
	return jpeg.Encode(writer, matrix.Image(), &jpeg.Options{Quality: quality})
}

// save a image matrix as a png,if unsuccessful it will return a error.
func saveAsPng(writer io.Writer, matrix Matrix) error {
	return png.Encode(writer, matrix.Image())
}

// save a image matrix as a gif,if unsuccessful it will return a error.
func saveAsGif(writer io.Writer, matrix []Matrix, delay []int, loopCount int) error {
	paletteds := make([]*image.Paletted, len(matrix))
	for i, m := range matrix {
		img := m.Image()
		paletteds[i] = image.NewPaletted(image.Rect(0, 0, m.Width(), m.Height()), palette.Plan9)
		draw.Draw(paletteds[i], paletteds[i].Bounds(), img, image.ZP, draw.Src) //添加图片
	}

	g := &gif.GIF{
		Image:     paletteds,
		Delay:     delay,
		LoopCount: loopCount,
	}
	return gif.EncodeAll(writer, g)
}
