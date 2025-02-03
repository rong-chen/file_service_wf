package menu

import (
	"file_service/api/authority"
	"file_service/api/user"
	"file_service/model/common/response"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		response.FailWithMessage("查询失败", c)
		return
	}
	users := user.ContextUser.FindUserInfo("id", id.(uint))
	authorities, err := authority.FindAuthorities(users.AuthorityId)
	if err != nil {
		response.FailWithMessage("查询失败"+err.Error(), c)
		return
	}

	var menuIds []uint
	for _, v := range authorities {
		menuIds = append(menuIds, v.MenuId)
	}

	list, err := FindMenuList(menuIds)
	if err != nil {
		response.FailWithMessage("查询失败"+err.Error(), c)
		return
	}
	response.OkWithData(list, "", c)
}
