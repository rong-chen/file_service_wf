package main

import (
	"file_service/core"
	"file_service/global"
	"file_service/initialize"
)

func main() {
	global.QY_VP = core.Viper()           // 初始化Viper
	global.QY_Db = initialize.GormMysql() // gorm连接数据库
	if global.QY_Db != nil {
		initialize.RegisterTables() // 初始化表
		initialize.InitDbData()     // 初始化表数据
	}
	defer func() {
		if global.QY_Db != nil {
			db, _ := global.QY_Db.DB()
			err := db.Close()
			if err != nil {
			}
		}
	}()
	core.RunWindowsServer()
}
