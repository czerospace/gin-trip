package main

import "github.com/gin-gonic/gin"

/*
	header 里加 Referer 最简单的作用就是防爬虫,请求 header 里不带这个 Referer 就不返回
*/
type Header struct {
	Referer string `header:"Referer" binding:"required"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		var h Header
		if err := c.ShouldBindHeader(&h); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Code":    0,
			"Referer": h.Referer,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
