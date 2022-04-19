package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// GET 获取所有文件信息
	r.GET("/posts", func(c *gin.Context) {
		c.String(http.StatusOK, "GET")
	})
	// POST 创建一篇文件
	r.POST("/posts", func(c *gin.Context) {
		c.String(http.StatusOK, "POST")
	})
	// PUT 修改一篇文章
	r.PUT("/posts/:id", func(c *gin.Context) {
		c.String(http.StatusOK, fmt.Sprintf("PUT id: %s", c.Param("id")))
	})
	// DELETE 删除一篇文章
	r.DELETE("/posts", func(c *gin.Context) {
		c.String(http.StatusOK, "DELETE")
	})

	// 匹配所有请求方法
	r.Any("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "any")
	})

	err := r.Run()
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
}
