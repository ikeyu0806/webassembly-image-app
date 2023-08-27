package main

import (
	"fmt"
	"syscall/js"
)

func getFileSize(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return nil
	}

	file := args[0]
	size := file.Get("size").Int()

	fmt.Printf("Uploaded File Size: %d bytes\n", size)

	return nil
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("getFileSize", js.FuncOf(getFileSize))

	<-c
}
