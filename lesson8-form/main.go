package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/form_post", postHandler)
	r.POST("/form_array", arrayHandler)
	r.POST("/form_map", mapHandler)
	err := r.Run()
	if err != nil {
		return
	}
}

func mapHandler(c *gin.Context) {
	user := c.PostFormMap("user")
	c.JSON(200, gin.H{
		"data": user,
	})
}

func arrayHandler(c *gin.Context) {
	ids := c.PostFormArray("ids")
	c.JSON(200, gin.H{
		"ids": ids,
	})
}

func postHandler(c *gin.Context) {
	message := c.PostForm("message")
	name := c.DefaultPostForm("name", "Niko")
	c.JSON(200, gin.H{
		"message": message,
		"name":    name,
	})
}
