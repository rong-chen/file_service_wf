package initialize

import (
	"file_service/global"
	"file_service/middleware"
	"file_service/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}
	return f, nil
}

// 初始化总路由
func Routers() *gin.Engine {
	Router := gin.New()
	Router.Use(gin.Recovery())
	// 配置 CORS：允许所有域名、所有方法、所有请求头
	Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许所有域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的 HTTP 方法
		AllowHeaders:     []string{"*"},                                       // 允许所有请求头
		AllowCredentials: true,                                                // 是否允许携带凭证
	}))
	Router.StaticFS(global.QY_CONFIG.Local.StorePath, justFilesFilesystem{http.Dir(global.QY_CONFIG.Local.StorePath)})
	// 方便统一添加路由组前缀 多服务器上线使用
	CheckGroup := Router.Group(global.QY_CONFIG.System.RouterPrefix)
	NoCheckGroup := Router.Group(global.QY_CONFIG.System.RouterPrefix)

	CheckGroup.Use(middleware.JWTAuth())
	{
		// 注册路由信息 第一公共的，不需要校验token
		for _, routers := range router.NoCheckRoutersList {
			routers.InitRouter(NoCheckGroup)
		}
		for _, routers := range router.CheckRoutersList {
			routers.InitRouter(CheckGroup)
		}
	}
	global.QY_ROUTERS = Router.Routes()
	return Router
}
