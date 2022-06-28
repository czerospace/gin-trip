package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"strings"
)

type User struct {
	// 需求： id 如果非1开头则验证失败，反之成功
	Id string `form:"id" binding:"required,niko"`
}

var customValidate validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().Interface().(string)
	if strings.HasPrefix(data, "1") {
		return true
	}
	return false
}

func main() {
	r := gin.Default()
	// 注册 customValidate
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("niko", customValidate)
	}
	r.GET("/user", func(c *gin.Context) {
		var u User
		if err := c.ShouldBind(&u); err != nil {
			c.JSON(200, gin.H{
				"Code": 200,
				"Msg":  err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"Code": 200,
			"Msg":  u.Id,
		})
	})

	err := r.Run()
	if err != nil {
		return
	}
}
