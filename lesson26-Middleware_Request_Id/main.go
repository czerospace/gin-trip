package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"log"
)

// RequestIdMiddleware 自定义中间件
func RequestIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Request.Header.Get("Niko")
		if id == "" {
			u, _ := uuid.NewV4()
			id = u.String()
			log.Println(id)
		}
		c.Writer.Header().Set("Niko", id)
		c.Next()
	}
}

func main() {
	r := gin.Default()
	r.Use(RequestIdMiddleware())
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
