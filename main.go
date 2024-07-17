// main.go
package main

import (
	"syscall/js"
)

func add(this js.Value, p []js.Value) interface{} {
    sum := p[0].Int() + p[1].Int()
    return js.ValueOf(sum)
}

func registerCallbacks() {
    js.Global().Set("add", js.FuncOf(add))
}

func main() {
    c := make(chan struct{}, 0)
    println("Wasm Go Initialized")
    registerCallbacks()
    <-c
}
