package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Header struct {
	Referer string `header:"Referer" binding:"required"`
}

/*
// 第一种写法
func middleware1(c *gin.Context) {

}

// 第二种写法 常用
func middleware2() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
*/

// RefererMiddleware 自定义一个反爬虫 refer
func RefererMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first in...")
		// 取到 referer
		ref := c.GetHeader("Referer")
		if ref == "" {
			c.AbortWithStatusJSON(200, gin.H{
				"msg": "非法访问",
			})
			return
		}
		c.Next()
		fmt.Println("first done...")
	}
}

func main() {
	r := gin.Default()
	//// 第一种写法使用
	//r.Use(middleware1)
	//// 第二种写法使用
	//r.Use(middleware2())
	r.Use(RefererMiddleware())
	// 第三种写法
	r.Use(func(c *gin.Context) {
		fmt.Println("second in ...")
		fmt.Println("我是第二个中间件")
		c.Next()
		fmt.Println("second done ...")
	})
	r.GET("ping", func(c *gin.Context) {
		fmt.Println("我是 core")
		c.JSON(200, gin.H{
			"code": 0,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}

}
