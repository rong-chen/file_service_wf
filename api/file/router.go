package file

import "github.com/gin-gonic/gin"

type Router struct {
}

func (Router) InitRouter(g *gin.RouterGroup) {
	router := g.Group("file")
	{
		//查询上传记录
		router.POST("findFile", FindFile)
		//上传片段
		router.POST("upload-chunk-file", UploadChunkFile)
		// 上传完成
		router.GET("finish", UploadSuccess)
		// 获取当前用户文件
		router.GET("find-file-list", FindFileList)
		// 获取所有文件
		router.GET("all-file-list", FindAllFileList)
		// 收藏文件
		router.POST("collection", Collection)
		// 获取下载key
		router.GET("download-key/:fileId", RegisterDownloadKey)
		// 下载文件
		router.GET("download/*key", DownLoadFile)
		router.DELETE("delete", Delete)
	}
}
