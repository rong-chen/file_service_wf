package userManage

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(g *gin.RouterGroup) {
	router := g.Group("user_v2")
	{
		// 用户列表
		router.GET("/list", List)
		//同意注册用户
		router.POST("/setting_user", ChangeUserInfo)
	}
}
