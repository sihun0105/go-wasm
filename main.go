package main

import "syscall/js"

func main() {
	println("Wasm loaded.")

	alert := js.Global().Get("alert")

	alert.Invoke("Hello Wasm!")
}

func Add(x int, y int) int {
	return x + y
}

// GOOS=js GOARCH=wasm go build -o main.wasm main.go
// cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .