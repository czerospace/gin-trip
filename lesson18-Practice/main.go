package main

import (
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	UserName   string `json:"username" binding:"required,min=3,max=7"`
	PassWord   string `json:"password" binding:"required,len=8"`
	RePassword string `json:"re_password" binding:"required,len=8"`
}

type RegisterForm struct {
	UserName string `json:"username" binding:"required,min=3,max=7"`
	PassWord string `json:"password" binding:"required,len=8"`
	Age      uint32 `json:"age" binding:"required,gte=1,lte=150"`
	Sex      uint32 `json:"sex" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

func main() {
	r := gin.Default()
	// 登陆
	r.POST("login", loginHandler)
	// 注册
	r.POST("register", registerHandler)
	err := r.Run()
	if err != nil {
		return
	}
}

// registerHandler 注册
func registerHandler(c *gin.Context) {
	var r RegisterForm
	if err := c.ShouldBindJSON(&r); err != nil {
		c.JSON(200, gin.H{
			"code": 40002,
			"msg":  "注册失败，请检查参数",
			"err":  err.Error(),
		})
		return
	}
	// 注册成功
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": r,
	})
}

// loginHandler 登陆
func loginHandler(c *gin.Context) {
	var l LoginForm
	if err := c.ShouldBindJSON(&l); err != nil {
		c.JSON(200, gin.H{
			"code": 40001,
			"err":  err.Error(),
			"msg":  "登陆失败，请检查参数",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "success",
		"data": l.UserName,
	})
}
