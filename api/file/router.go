package file

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(g *gin.RouterGroup) {
	router := g.Group("file")
	{
		router.POST("findFile", FindFile)
		router.POST("upload-chunk-file", UploadChunkFile)
		router.GET("finish", UploadSuccess)
		router.GET("find-file-list", FindFileList)
		router.GET("all-file-list", FindAllFileList)
	}
}
