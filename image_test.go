package image_lib

import (
	"testing"
)

func TestMustRead(t *testing.T) {
	imgMatrix, typeName, err := OpenImage("./test/test.jpg")
	if err != nil {
		panic(err)
	}
	if typeName != "jpeg" {
		t.Error("type error")
	}
	//imgMatrix2, typeName, err := OpenImage("./test/tt.jpg")
	//ToBinaryImage(imgMatrix)
	im := TOGray(imgMatrix)
	err = Save2File("test/tt.jpg", []Matrix{im}, typeName, 100)
	if err != nil {
		panic(err)
	}
}
