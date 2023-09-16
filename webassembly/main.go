package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"syscall/js"
)

func getFileSize(this js.Value, args []js.Value) interface{} {
	imageData := args[0]
	uint8Array := js.Global().Get("Uint8Array").New(imageData)
	imageBytes := make([]byte, uint8Array.Length())
	js.CopyBytesToGo(imageBytes, uint8Array)

	fmt.Printf("Uploaded File Size: %d bytes\n", len(imageBytes))

	return nil
}

func convertToBlackAndWhite(this js.Value, args []js.Value) interface{} {
	imageData := args[0]
	uint8Array := js.Global().Get("Uint8Array").New(imageData)
	imageBytes := make([]byte, uint8Array.Length())
	js.CopyBytesToGo(imageBytes, uint8Array)

	format := "jpeg"

	var img image.Image
	var err error
	switch format {
	case "jpeg":
		img, err = jpeg.Decode(bytes.NewReader(imageBytes))
	case "png":
		img, err = png.Decode(bytes.NewReader(imageBytes))
	default:
		fmt.Println("サポートされていない画像フォーマット:", format)
		return nil
	}

	if err != nil {
		fmt.Println("画像デコードエラー:", err)
		return nil
	}
	fmt.Printf("img: %v", img)

	// resultArray := js.Global().Get("Uint8Array").New(len(imageBytes))
	// js.CopyBytesToJS(resultArray, imageBytes)

	return nil
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("getFileSize", js.FuncOf(getFileSize))
	js.Global().Set("convertToBlackAndWhite", js.FuncOf(convertToBlackAndWhite))

	<-c
}
