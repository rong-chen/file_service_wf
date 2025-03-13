package userManage

import (
	"file_service/model/common/response"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	authorityId, ok := c.Get("authorityId")
	if !ok {
		response.FailWithMessage("用户权限获取错误", c)
		return
	}
	uAuthorityId := authorityId.(uint)
	if uAuthorityId != 888 {
		response.FailWithMessage("暂无权限", c)
		return
	}
	list := FindAllUserInfo()
	for i, _ := range list {
		list[i].Password = ""
	}
	response.OkWithData(map[string]interface{}{
		"list": list,
	}, "获取成功", c)
}

func ChangeUserInfo(c *gin.Context) {
	authorityId, ok := c.Get("authorityId")
	if !ok {
		response.FailWithMessage("用户权限获取错误", c)
		return
	}
	uAuthorityId := authorityId.(uint)
	if uAuthorityId != 888 {
		response.FailWithMessage("暂无权限", c)
		return
	}

	var p UpdateParams
	err := c.ShouldBindJSON(&p)
	if err != nil {
		response.FailWithMessage("变更失败"+err.Error(), c)
		return

	}
	if p.IsExamine != true && p.IsExamine != false && p.MountPath != "" {
		response.FailWithMessage("参数错误", c)
	}

	err = UpdateUsers(p)
	if err != nil {
		response.FailWithMessage("变更失败"+err.Error(), c)
		return
	}
	response.OkWithMessage("获取成功", c)
}
