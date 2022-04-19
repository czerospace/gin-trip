package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载静态文件
	r.Static("/images", "./images")
	// StaticFS works just like `Static()` but a custom `http.FileSystem` can be used instead.
	r.StaticFS("/static", http.Dir("./static"))
	// 加载静态文件
	r.StaticFile("code", "main.go")
	err := r.Run()
	if err != nil {
		return
	}
}
