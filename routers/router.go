package routers

import (
	"anroid/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},              // 允许所有来源
		AllowMethods:     []string{"*"},              // 允许的请求方法
		AllowHeaders:     []string{"*"},              // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"}, // 允许暴露的响应头
		AllowCredentials: true,                       // 允许携带凭证（如Cookie）
	}))
	V1Group := r.Group("api")
	{
		//用户注册
		V1Group.POST("/user/register", controller.UserRegister)
		//用户登录
		V1Group.POST("/user/login", controller.UserLogin)
		//申请验证码
		V1Group.POST("/user/register-email", controller.SendEmailRegister)
		//返回所有诗句
		V1Group.GET("/user/poetry", controller.SendPoetry)
		//返回所有评论
		V1Group.GET("/user/comment", controller.SendComment)
		//返回所有诗作
		V1Group.GET("/user/poetical", controller.SendPoetical)
		//返回所有签名
		V1Group.GET("/user/signature", controller.SendSignature)
		//添加评论
		V1Group.POST("/user/put/comment", controller.CreatComments)
		//笔记相关接口
		V1Group.POST("/note/create", controller.CreateNote)
		V1Group.GET("/note/list", controller.GetNotes)
		V1Group.PUT("/note/update", controller.UpdateNote)
		V1Group.DELETE("/note/delete", controller.DeleteNote)
		//图片上传
		V1Group.POST("/upload/image", controller.UploadImage)
		//替换头像
		V1Group.PUT("/user/avatar", controller.UpdateAvatar)
	}
	return r
}
