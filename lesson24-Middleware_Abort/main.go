package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func middleware1(c *gin.Context) {
	log.Println("middleware1 in ...")
	c.Set("key", 1000)
	log.Println("middleware1 before next ...")
	// 判断
	k := c.GetInt("key")
	if k == 1000 {
		// c.Abort()
		c.AbortWithStatusJSON(200, gin.H{
			"msg": "验证失败",
		})
		return
	}
	c.Next()
	log.Println("middleware1 next after ...")
	log.Println("middleware1 done")

}

func middleware2(c *gin.Context) {
	log.Println("middleware2 in ...")
	log.Println("middleware2 before next ...")
	c.Next()
	log.Println("middleware2 next after ...")
	log.Println("middleware2 done")
}

func main() {
	r := gin.Default()
	r.Use(middleware1, middleware2)
	r.GET("ping", func(c *gin.Context) {
		log.Println("func in core ...")
		k := c.GetInt("key") // 1000
		c.Set("key", k+2000)
		log.Println("func in core done")
		c.JSON(200, gin.H{
			"msg": c.GetInt("key"),
		})
	})
	err := r.Run()
	if err != nil {
		return
	}
}
