package file_v2

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(g *gin.RouterGroup) {
	router := g.Group("file_v2")
	{
		router.POST("check_file", CheckFile)
	}
}
