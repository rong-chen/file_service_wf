package router

import (
	"file_service/api/file"
	"file_service/api/file_collection"
	"file_service/api/group_share"
	"file_service/api/menu"
	"file_service/api/user"
	"file_service/api/v2/file_v2"
	"file_service/api/v2/userManage"
	"github.com/gin-gonic/gin"
)

type Routers interface {
	InitRouter(group *gin.RouterGroup)
}

var NoCheckRoutersList = []Routers{
	&user.UserRouter{},
}
var CheckRoutersList = []Routers{
	&user.UserInfoRouter{},
	&file.Router{},
	&file_collection.Router{},
	&menu.Router{},
	&group_share.Router{},
	&file_v2.Router{},
	&userManage.Router{},
}
