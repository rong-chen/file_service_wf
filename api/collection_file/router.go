package collection_file

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func (*Router) InitRouter(group *gin.RouterGroup) {
	r := group.Group("/file")
	{

		r.GET("like-list", List)
		r.POST("like", Like)
		r.GET("find/*val", FindMusic)
	}
}
