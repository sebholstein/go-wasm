package main

import (
	"syscall/js"
	"time"
)

func main() {
	js.Global().Get("console").Call("log", "Hello world Go/wasm!")
	js.Global().Get("document").Call("getElementById", "app").Set("innerText", time.Now().String())
}
