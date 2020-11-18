package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("hello", js.FuncOf(hello))
	global.Set("changeContent", js.FuncOf(changeContent))

	<-done
}

func changeContent(this js.Value, args []js.Value) interface{} {
	fmt.Println("Call change content")
	document := js.Global().Get("document")
	content := document.Call("getElementById", "content")

	content.Set("textContent", fmt.Sprintf("%v Hello WebAssembly!", content.Get("textContent")))
	return nil
}

func hello(this js.Value, args []js.Value) interface{} {
	fmt.Println("Hello WebAssembly")
	js.Global().Call("alert", "Hello WebAssembly")
	return nil
}
