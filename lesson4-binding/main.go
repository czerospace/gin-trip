package main

import "github.com/gin-gonic/gin"

// 泛绑定

func main() {
	r := gin.Default()
	// 非泛绑定
	r.GET("/usr/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})

	// 泛绑定
	r.GET("/posts/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.JSON(200, gin.H{
			"name":   name,
			"action": action,
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
