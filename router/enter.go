package router

import (
	"file_service/api/collection_file"
	"file_service/api/file"
	"file_service/api/menu"
	"file_service/api/user"
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
	&collection_file.Router{},
	&menu.Router{},
}
