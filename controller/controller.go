package controller

import (
	"anroid/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户注册并返回信息给客户端

func UserRegister(c *gin.Context) {
	var requestData struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
		Email    string `form:"email" json:"email"`
		Code     string `form:"code" json:"code"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 调用UserRegister函数进行注册
	result := server.UserRegister(requestData.Username, requestData.Password, requestData.Email, requestData.Code)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "注册成功",
		})

	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}

}

// 用户登录并返回信息给客户端

func UserLogin(c *gin.Context) {
	var requestData struct {
		Username string `form:"username" json:"username"`
		Password string `form:"password" json:"password"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	result := server.UserLogint(requestData.Username, requestData.Password)
	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "登录成功",
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Code":    http.StatusNotFound,
			"Message": result,
		})
	}

}

// 请求验证码

func SendEmailRegister(c *gin.Context) {
	var requestData struct {
		Email string `form:"email" json:"email"`
	}
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}
	// 调用UserRegister函数进行注册
	result := server.SendEmail(requestData.Email, false)

	// 根据注册结果返回相应的数据给前端

	if result == "" {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": "邮件已发送，请注意查收",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Code":    http.StatusOK,
			"Message": result,
		})
	}

}
