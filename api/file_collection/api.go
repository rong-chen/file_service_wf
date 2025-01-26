package file_collection

import (
	"file_service/model/common/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func Like(c *gin.Context) {
	type Params struct {
		FileId uint `json:"file_id"`
	}
	var params Params
	userId, _ := c.Get("user_id")
	err := c.BindJSON(&params)
	var lf LikeFile
	lf.FileId = params.FileId
	lf.UserId = userId.(uint)
	if err != nil {
		response.FailWithMessage("参数错误:"+err.Error(), c)
		return
	}
	err = CreateLikeFile(lf)
	if err != nil {
		response.FailWithMessage("", c)
		return
	}
	response.OkWithMessage("操作成功", c)
}

func List(c *gin.Context) {
	userId, _ := c.Get("user_id")
	files, err := FindAllListByUserId(userId.(uint))
	if err != nil {
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(files, "", c)
}

func FindMusic(c *gin.Context) {
	name := c.Param("val")
	name = strings.TrimPrefix(name, "/")
	val, err := FindMusicListByFileVal(name)
	if err != nil {
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithData(val, "", c)
}
