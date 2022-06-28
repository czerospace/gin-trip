package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
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

var trans ut.Translator

func main() {
	if err := InitializaTrans(); err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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
		// 判断是不是 validator 错误
		err, ok := err.(validator.ValidationErrors)
		// 如果不是 validator 错误，返回 注册失败，请检查参数
		if !ok {
			c.JSON(200, gin.H{
				"code": 40010,
				"msg":  "注册失败，请检查参数",
				"err":  err.Error(),
			})
			return
		}
		// 如果是 validator 错误，翻译成中文
		c.JSON(200, gin.H{
			"code": 40004,
			"err":  err.Translate(trans),
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

// InitializaTrans 翻译函数
func InitializaTrans() (err error) {
	// 修改 gin 框架 validator 引擎属性
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		zh := zh.New()
		uni := ut.New(zh, zh)
		trans, _ = uni.GetTranslator("zh")
		err = zh_translations.RegisterDefaultTranslations(v, trans)
		return
	}
	return
}
