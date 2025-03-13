package initialize

import (
	"file_service/api/authority"
	"file_service/api/file"
	"file_service/api/file_collection"
	"file_service/api/group_share"
	"file_service/api/menu"
	"file_service/api/user"
	"file_service/api/v2/file_v2"
	"file_service/config"
	"file_service/global"
	"fmt"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct{}

type Writer struct {
	config config.GeneralDB
	writer logger.Writer
}

// Printf 格式化打印日志
func (c *Writer) Printf(message string, data ...any) {
	if c.config.LogZap {
		switch c.config.LogLevel() {
		case logger.Silent:
			global.QY_LOG.Debug(fmt.Sprintf(message, data...))
		case logger.Error:
			global.QY_LOG.Error(fmt.Sprintf(message, data...))
		case logger.Warn:
			global.QY_LOG.Warn(fmt.Sprintf(message, data...))
		case logger.Info:
			global.QY_LOG.Info(fmt.Sprintf(message, data...))
		default:
			global.QY_LOG.Info(fmt.Sprintf(message, data...))
		}
		return
	}
}

func NewWriter(config config.GeneralDB) *Writer {
	return &Writer{config: config}
}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {

	return &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // 输出到标准输出
			logger.Config{
				SlowThreshold: 200 * time.Millisecond,
				LogLevel:      logger.Info,
				Colorful:      true,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
}

func GormMysql() *gorm.DB {
	m := global.QY_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic("数据库连接失败")
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func GormRedis() *redis.Client {
	rc := global.QY_CONFIG.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     rc.Addr,
		Password: rc.Password, // no password set
		DB:       0,           // use default DB
	})
	return rdb
}

func RegisterTables() {
	db := global.QY_Db
	err := db.AutoMigrate(
		user.Users{},
		authority.Authorities{},
		file.FileChunk{},
		file.File{},
		file.ShareFileInfo{},
		file_collection.LikeFile{},
		menu.BaseMenu{},
		authority.AuthoritiesMenu{},
		group_share.Group{},
		group_share.GroupUsers{},
		group_share.GroupFiles{},
		file_v2.FileInfo{},
		file_v2.ChunkInfo{},
	)
	if err != nil {
		os.Exit(0)
	}
}

func InitDbData() {
	err := menu.InitMenuDbData()
	if err != nil {
		return
	}
	err = user.InitUserDbData()
	if err != nil {
		return
	}
	err = authority.InitAuthoritiesData()
	if err != nil {
		return
	}
}
