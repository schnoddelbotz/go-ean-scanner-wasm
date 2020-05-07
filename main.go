package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	js.Global().Set("suggest", js.FuncOf(suggest))
	<-c
}

// define JS response struct ...? return / set it, display!
type NewUser struct {
	Email    string
	Password string
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
