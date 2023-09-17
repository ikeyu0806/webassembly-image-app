package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
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

func showUploadedImage(this js.Value, args []js.Value) interface{} {
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

func convertToBlackAndWhite(this js.Value, args []js.Value) interface{} {
	imageData := args[0]
	uint8Array := js.Global().Get("Uint8Array").New(imageData)
	imageBytes := make([]byte, uint8Array.Length())
	js.CopyBytesToGo(imageBytes, uint8Array)

	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return err.Error()
	}

	grayscaleImg := convertToGrayscale(img)

	var encodedImage bytes.Buffer
	err = png.Encode(&encodedImage, grayscaleImg)
	if err != nil {
		return err.Error()
	}

	dataURI := "data:image/png;base64," + base64.StdEncoding.EncodeToString(encodedImage.Bytes())

	imgElement := js.Global().Get("document").Call("createElement", "img")
	imgElement.Set("src", dataURI)

	var body = js.Global().Get("document").Get("body")
	body.Call("appendChild", imgElement)

	return nil
}

func convertToGrayscale(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := src.At(x, y).RGBA()
			gray := uint8((r + g + b) / 3)
			dst.SetGray(x, y, color.Gray{Y: gray})
		}
	}
	return dst
}

func convertToSepia(this js.Value, args []js.Value) interface{} {
	imageData := args[0]
	uint8Array := js.Global().Get("Uint8Array").New(imageData)
	imageBytes := make([]byte, uint8Array.Length())
	js.CopyBytesToGo(imageBytes, uint8Array)

	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		return err.Error()
	}

	sepiaImg := convertToSepiaImage(img)

	var encodedImage bytes.Buffer
	err = png.Encode(&encodedImage, sepiaImg)
	if err != nil {
		return err.Error()
	}

	dataURI := "data:image/png;base64," + base64.StdEncoding.EncodeToString(encodedImage.Bytes())

	imgElement := js.Global().Get("document").Call("createElement", "img")
	imgElement.Set("src", dataURI)

	var body = js.Global().Get("document").Get("body")
	body.Call("appendChild", imgElement)

	return nil
}

func convertToSepiaImage(src image.Image) image.Image {
	bounds := src.Bounds()
	sepiaImg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := src.At(x, y)
			originalRGBA := color.RGBAModel.Convert(originalColor).(color.RGBA)

			// 赤、緑、青の合計を3で割ることでグレースケールの輝度値を取得
			gray := (uint32(originalRGBA.R) + uint32(originalRGBA.G) + uint32(originalRGBA.B)) / 3
			// グレースケールの輝度値にセピア調のフィルター適用
			sepiaR := gray + 112
			sepiaG := gray + 66
			sepiaB := gray + 20

			if sepiaR > 255 {
				sepiaR = 255
			}
			if sepiaG > 255 {
				sepiaG = 255
			}
			if sepiaB > 255 {
				sepiaB = 255
			}

			sepiaImg.Set(x, y, color.RGBA{R: uint8(sepiaR), G: uint8(sepiaG), B: uint8(sepiaB), A: originalRGBA.A})
		}
	}

	return sepiaImg
}

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("getFileSize", js.FuncOf(getFileSize))
	js.Global().Set("showUploadedImage", js.FuncOf(showUploadedImage))
	js.Global().Set("convertToBlackAndWhite", js.FuncOf(convertToBlackAndWhite))
	js.Global().Set("convertToSepia", js.FuncOf(convertToSepia))

	<-c
}
