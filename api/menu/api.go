package menu

import (
	"file_service/model/common/response"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	list, err := FindMenuList()
	if err != nil {
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(list, "", c)
}
