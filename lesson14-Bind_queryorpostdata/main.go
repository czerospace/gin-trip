package main

import "github.com/gin-gonic/gin"

/*
	ShouldBind 可以从 content-type 里获取数据类型，然后去绑定具体的数据类型
	但是一般是啥格式就直接绑定啥格式,比如直接 ShouldBindJSON，生产环境别用 ShouldBind
*/

type User struct {
	ID string `json:"id" binding:"required"`
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
			"Msg":  user.ID,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
