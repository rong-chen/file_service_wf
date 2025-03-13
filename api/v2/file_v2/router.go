package file_v2

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(g *gin.RouterGroup) {
	router := g.Group("file_v2")
	{
		router.POST("check_file", CheckFile)
		router.POST("upload-chunk", UploadChunk)
		router.POST("combined-file", CombinedFile)
		router.POST("list", List)
		//删除文件
		router.DELETE("delete", Delete)
	}
}
