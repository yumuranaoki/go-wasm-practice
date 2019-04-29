package main

import (
	"syscall/js"
)

func increment(this js.Value, values []js.Value) interface{} {
	document := js.Global().Get("document")
	root := document.Call("getElementById", "root")
	root.Set("innerText", "incremented")
	return nil
}

func main() {
	c := make(chan bool)
	js.Global().Set("increment", js.FuncOf(increment))
	<-c
}
