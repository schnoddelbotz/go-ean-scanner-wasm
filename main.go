package main

import (
	"fmt"
	"image"
	"syscall/js"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"

	"github.com/schnoddelbotz/suggest-wasm/filter"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	c := make(chan struct{}, 0)

	println("WASM Go Initializing...")

	// register suggest func here to run when JS suggest function is called
	js.Global().Set("suggest", js.FuncOf(suggest))

	// webcam filter
	applyFilterFunc := js.FuncOf(applyFilter) // wrapper function
	js.Global().Set("applyFilter", applyFilterFunc)
	defer applyFilterFunc.Release()

	<-c
}

var uint8Array = js.Global().Get("Uint8Array")
var threshold = 127

func applyFilter(this js.Value, args []js.Value) interface{} {

	val := args[0].Int()
	length := args[1].Length()

	jsPixels := make(filter.Pixels, length)

	_ = js.CopyBytesToGo(jsPixels, args[1])

	if val == 0 {
		jsPixels.MakeGrey()
	} else if val == 1 {
		jsPixels.Invert()
	} else if val == 2 {
		jsPixels.MakeNoise()
	} else if val == 3 {
		jsPixels.MakeRed()
	}

	buf := uint8Array.New(len(jsPixels))
	js.CopyBytesToJS(buf, jsPixels)

	// this is super inefficient, drops 60 to 20 FPS on me Mac.
	// avoid copying if possible at all. also, drop filter stuff
	// above if we do not want to apply image fun (which we might rather
	// want to do in CSS anyway...).
	reader := oned.NewMultiFormatUPCEANReader(nil)
	r := image.Rect(0,0, 640, 480)
	img := image.NewRGBA(r)
	_ = js.CopyBytesToGo(img.Pix, args[1])
	b, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		println("err", err)
	}
	ean, derr := reader.DecodeWithoutHints(b)
	if derr == nil {
		fmt.Printf("We found EAN: %s\n", ean.String())
		js.Global().Set("eanFound", true)
		js.Global().Set("ean", ean.String())
	}

	// can we return static error image here in case of error? :-)
	return buf
}

func suggest(t js.Value, tt []js.Value) interface{} {
	query := tt[0]
	// this goes to browser console log
	fmt.Printf("suggest called with q=%s\n", query)
	// fixme legacy way?
	//js.Global().Set("output", js.ValueOf("foo32"))
	// fixme new?
	//user := NewUser{Email: "Joe@Foo.com", Password: "Jim2000"}
	//return user
	data := `{"foo":"bar"}`
	return data
}
