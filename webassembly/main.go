package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("fmt Println!")
	js.Global().Get("console").Call("log", "js.Global Call!")
	select {}
}
