package main

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"syscall/js"
)

func loadImageData(file js.Value) (image.Image, error) {
	arrayBuffer := file.Get("arrayBuffer")
	buffer := make([]byte, arrayBuffer.Get("byteLength").Int())
	js.CopyBytesToGo(buffer, arrayBuffer)

	imageReader := bytes.NewReader(buffer)
	img, _, err := image.Decode(imageReader)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func getFileSize(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return nil
	}

	file := args[0]
	size := file.Get("size").Int()

	fmt.Printf("Uploaded File Size: %d bytes\n", size)

	img, err := loadImageData(file)
	if err != nil {
		fmt.Printf("Error loading image: %v\n", err)
		return nil
	}

	fmt.Printf("img result%v", img)

	return nil
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("getFileSize", js.FuncOf(getFileSize))

	<-c
}
