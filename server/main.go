package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/js", "./webassembly/js")
	r.Static("/wasm", "./webassembly")
	r.StaticFile("/", "./index.html")

	r.Run(":80")
}
