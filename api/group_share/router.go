package group_share

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(g *gin.RouterGroup) {
	router := g.Group("group")
	{
		router.POST("add", CreateGroup)
		router.DELETE("del", DeleteGroup)
		router.GET("list", GetGroup)
	}
	router2 := g.Group("group-user")
	{
		router2.POST("join", Join)
		router2.GET("list", FindGroupUsersList)
	}

	router3 := g.Group("group-file")
	{
		router3.POST("add", AddFile)
		router3.GET("list", FindGroupFilesList)
	}
}
