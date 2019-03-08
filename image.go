package image_lib

import (
	"errors"
	"fmt"
	"image"
	"os"

	_ "image/gif"
	_ "image/png"
)

func OpenImage(filePath string) (matrix Matrix, typeName string, err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	var img image.Image

	img, typeName, err = image.Decode(f)

	bounds := img.Bounds()
	matrix = &NRGBAMatrix{}
	matrix.Matrix(img, bounds.Max.X, bounds.Max.Y)
	return
}

func Save2File(filePath string, matrixs []Matrix, typeName string, options ...interface{}) error {
	outfile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outfile.Close()

	switch typeName {
	case "jpeg":
		return saveAsJPEG(outfile, matrixs[0], options[0].(int))
	case "png":
		return saveAsPng(outfile, matrixs[0])
	case "gif":
		return saveAsGif(outfile, matrixs, options[0].([]int), options[1].(int))
	default:
		return errors.New("do not support this image type")
	}
}

func ToBinaryImage(matrix Matrix) {
	m := matrix.(*NRGBAMatrix)
	for i := 0; i < matrix.Height(); i++ {
		for j := 0; j < matrix.Width(); j++ {
			if (m.matrix[i][j][0]+m.matrix[i][j][1]+m.matrix[i][j][2])/3 > 125 {
				m.matrix[i][j][0] = 255
				m.matrix[i][j][1] = 255
				m.matrix[i][j][2] = 255
			} else {
				m.matrix[i][j][0] = 0
				m.matrix[i][j][1] = 0
				m.matrix[i][j][2] = 0
			}
		}
	}
}

func ToRGB(matrix Matrix) [][][3]uint8 {
	m := matrix.(*NRGBAMatrix)
	r := make([][][3]uint8, matrix.Height())
	for i := 0; i < matrix.Height(); i++ {
		r[i] = make([][3]uint8, matrix.Width())
		for j := 0; j < matrix.Width(); j++ {
			t := (255 - m.matrix[i][j][3]) / 255.0
			r[i][j][0] *= t
			r[i][j][1] *= t
			r[i][j][2] *= t
		}
	}
	return r
}

func TOGray(matrix Matrix) Matrix {
	m := matrix.(*NRGBAMatrix)
	r := make([][]uint8, matrix.Height())
	for i := 0; i < matrix.Height(); i++ {
		r[i] = make([]uint8, matrix.Width())
		for j := 0; j < matrix.Width(); j++ {
			r[i][j] = uint8(float32(m.matrix[i][j][0])*0.299 + float32(m.matrix[i][j][1])*0.587 + float32(m.matrix[i][j][2])*0.114)
			fmt.Println(r[i][j])
		}
	}
	return &GaryMatrix{r}
}
