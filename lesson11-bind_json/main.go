package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   string `json:"id" binding:"required,email"`
	Name string `json:"name" binding:"required,min=3,max=8"`
}

func main() {
	r := gin.Default()
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"Code": 0,
			"ID":   user.ID,
			"Name": user.Name,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
