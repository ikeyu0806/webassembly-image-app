package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.StaticFile("/", "./index.html")

	r.Run(":80")
}
