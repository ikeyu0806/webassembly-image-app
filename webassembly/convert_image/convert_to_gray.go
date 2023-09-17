package convert_image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"syscall/js"
)

func ConvertToBlackAndWhite(this js.Value, args []js.Value) interface{} {
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
