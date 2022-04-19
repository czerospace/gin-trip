package main

import "github.com/gin-gonic/gin"

// 路由分组

func main() {
	r := gin.Default()

	// 路由分组
	/*
		不分组写法
		r.GET("/posts", GetHandler)
		r.POST("/posts", PostHandler)
		// 删除 id=1 这篇文章
		r.DELETE("/posts/1", DeleteHandler)
	*/

	p := r.Group("/posts")
	{
		// http://127.0.0.1:8080/posts
		// Group 已经把 posts 加入到 URL 里面了，所以下面的 GET POST 和 DELETE 的 relativePath 里就不用再加 posts了
		p.GET("", GetHandler)
		p.POST("", PostHandler)
		// 删除 id=1 这篇文章
		p.DELETE("/:id", DeleteHandler)
	}

	// Simple group: v1
	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	v2 := r.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	// 嵌套
	api := r.Group("/api")
	{
		v3 := api.Group("posts")
		{
			v3.GET("", GetHandler)
			v3.POST("", PostHandler)
			// 删除 id=1 这篇文章
			v3.DELETE("/:id", DeleteHandler)
		}
	}

	err := r.Run()
	if err != nil {
		return
	}
}

func submitEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "submitEndpoint",
	})
}

func readEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "readEndpoint",
	})
}

func loginEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "loginEndpoint",
	})
}

func DeleteHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "DELETE",
	})
}

func PostHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "POST",
	})
}

func GetHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET",
	})
}
