package main

import (
	"syscall/js"

	"go-webassembly/webassembly/analyze_image"
	"go-webassembly/webassembly/convert_image"
)

func main() {
	c := make(chan struct{}, 0)

	js.Global().Set("getFileSize", js.FuncOf(analyze_image.GetFileSize))
	js.Global().Set("showUploadedImage", js.FuncOf(analyze_image.ShowUploadedImage))
	js.Global().Set("convertToBlackAndWhite", js.FuncOf(convert_image.ConvertToBlackAndWhite))
	js.Global().Set("convertToSepia", js.FuncOf(convert_image.ConvertToSepia))

	<-c
}
