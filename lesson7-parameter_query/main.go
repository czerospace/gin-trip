package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	// query string
	r.GET("welcome", welcomeHandler)
	// query array
	r.GET("array", arrayHandler)
	// query map
	r.GET("map", mapHandler)
	err := r.Run()
	if err != nil {
		return
	}
}

func mapHandler(c *gin.Context) {
	m := c.QueryMap("user")
	c.JSON(200, gin.H{
		"data": m,
	})
}

func arrayHandler(c *gin.Context) {
	ids := c.QueryArray("ids")
	c.JSON(200, gin.H{
		"ids": ids,
	})
}

func welcomeHandler(c *gin.Context) {
	firstName := c.DefaultQuery("firstname", "Niko")
	lastName := c.Query("lastname")
	c.JSON(200, gin.H{
		"firstname": firstName,
		"lastname":  lastName,
	})
}
