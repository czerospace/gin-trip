package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func LoginAuth(c *gin.Context) {
	fmt.Println("我是登陆保护中间件,login")
}

func main() {
	r := gin.Default()

	// 全局调用中间件
	//r.Use(LoginAuth)

	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	r.POST("login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "token",
		})
	})
	r.POST("register", func(c *gin.Context) {

	})
	user := r.Group("user", LoginAuth)
	{
		user.GET(":id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "我获取用户详情接口，需要登陆保护",
			})
		})
		user.PUT(":id", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "更新用户详情接口，需要登陆保护",
			})
		})
	}

	err := r.Run()
	if err != nil {
		return
	}
}
