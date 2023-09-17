package analyze_image

import (
	"encoding/base64"
	"fmt"
	"syscall/js"
)

func GetFileSize(this js.Value, args []js.Value) interface{} {
	imageData := args[0]
	uint8Array := js.Global().Get("Uint8Array").New(imageData)
	imageBytes := make([]byte, uint8Array.Length())
	js.CopyBytesToGo(imageBytes, uint8Array)

	fmt.Printf("Uploaded File Size: %d bytes\n", len(imageBytes))

	return nil
}

func ShowUploadedImage(this js.Value, args []js.Value) interface{} {
	imageData := args[0]
	uint8Array := js.Global().Get("Uint8Array").New(imageData)
	imageBytes := make([]byte, uint8Array.Length())
	js.CopyBytesToGo(imageBytes, uint8Array)

	newUint8Array := js.Global().Get("Uint8Array").New(len(imageBytes))
	js.CopyBytesToJS(newUint8Array, imageBytes)

	imgElement := js.Global().Get("document").Call("createElement", "img")

	dataURI := "data:image/png;base64," + base64.StdEncoding.EncodeToString(imageBytes)

	imgElement.Set("src", dataURI)

	var body = js.Global().Get("document").Get("body")
	body.Call("appendChild", imgElement)

	return nil
}
