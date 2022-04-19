package main

import (
	"github.com/gin-gonic/gin"
)

// 获取参数

// Person 创建一个结构体，绑定 数据类型
type Person struct {
	ID   int    `uri:"id"`
	Name string `uri:"name"`
}

func main() {
	r := gin.Default()

	/*
			访问 http://127.0.0.1:8080/9527/delete
			返回
				{
		  			"action": "delete",
		  			"id": "9527"
				}
	*/
	// id没有绑定类型，既可以传 int 也可以传 string
	/*
		r.GET("/:id/:action", func(c *gin.Context) {
			// 获取 id
			id := c.Param("id")
			// 获取action
			action := c.Param("action")
			c.JSON(200, gin.H{
				"id":     id,
				"action": action,
			})
		})
	*/

	// id 绑定 int 类型,只能传 int
	// 先创建结构体，在结构体里指定数据类型
	r.GET("/:id/:name", func(c *gin.Context) {
		var p Person
		if err := c.ShouldBindUri(&p); err != nil {
			c.Status(404)
			return
		}
		c.JSON(200, gin.H{
			"id":   p.ID,
			"name": p.Name,
		})

	})
	err := r.Run()
	if err != nil {
		return
	}
}
