package user

import (
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Group("user")
	{
		router.POST("/register", RegisterUser)
		router.POST("/login", Login)
	}
}

type UserInfoRouter struct{}

func (u *UserInfoRouter) InitRouter(Router *gin.RouterGroup) {
	router := Router.Group("user")
	{
		router.GET("/info", GetUserInfo)
		//同意注册用户
		router.POST("/consent", ConsentRegister)
		// 用户列表
		router.GET("/list", List)
	}
}
