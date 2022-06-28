package main

import "github.com/gin-gonic/gin"

/*
	ShouldBindQuery 如果数据校验失败，返回200
	(MUST)BindQuery 如果数据校验失败，返回400
*/
type User struct {
	// 9a69c5f3-6235-3d36-b27d-a3879d2b8f57
	ID string `form:"id" binding:"required,uuid"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindQuery(&user); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Code": 0,
			"Id":   user.ID,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
