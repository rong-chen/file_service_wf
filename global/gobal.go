package global

import (
	"file_service/config"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	QY_Db                  *gorm.DB
	QY_Redis               *redis.Client
	QY_CONFIG              config.Server
	QY_VP                  *viper.Viper
	QY_LOG                 *zap.Logger
	QY_Concurrency_Control = &singleflight.Group{}
	QY_ROUTERS             gin.RoutesInfo
)
