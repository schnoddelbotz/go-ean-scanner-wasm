package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)


func main() {
	// open and decode image file
	file, _ := os.Open("ean_ok.jpg")
	img, _, _ := image.Decode(file)

	// prepare BinaryBitmap
	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	reader := oned.NewMultiFormatUPCEANReader(nil)
	result, err := reader.Decode(bmp, nil)

	fmt.Println(result)
	fmt.Println(err)
}