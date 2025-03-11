package file_v2

import (
	"file_service/model/common/response"
	"github.com/gin-gonic/gin"
)

// CheckFile 该方法创建file，并且检测file是否已经上传了
func CheckFile(c *gin.Context) {
	id, _ := c.Get("user_id")
	var fileInfo FileInfo
	err := c.ShouldBindJSON(&fileInfo)
	if err != nil {
		response.FailWithMessage("参数错误:"+err.Error(), c)
		return
	}
	fileInfo.UserId = id.(uint)
	_, err = CreateOrFindFileInfo(fileInfo)
	if err != nil {
		response.FailWithMessage("检测文件错误:"+err.Error(), c)
		return
	}
	response.OkWithMessage("检测文件成功", c)
}
