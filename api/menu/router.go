package menu

import "github.com/gin-gonic/gin"

type Router struct {
}

func (*Router) InitRouter(group *gin.RouterGroup) {
	r := group.Group("/menu")
	{
		r.GET("list", List)
	}
}
