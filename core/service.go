package core

import (
	"file_service/global"
	"file_service/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	address := fmt.Sprintf(":%d", global.QY_CONFIG.System.Addr)
	Router := initialize.Routers()
	fmt.Println("项目启动：：：" + address)
	s := initServer(address, Router)
	fmt.Println(s.ListenAndServe().Error())
	// global.QY_LOG.Info("server run success on ", zap.String("address", address))
	// global.QY_LOG.Error()
}
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
}
