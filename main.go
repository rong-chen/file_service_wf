package main

import (
	"file_service/core"
	"file_service/global"
	"file_service/initialize"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	global.QY_VP = core.Viper()           // 初始化Viper
	global.QY_Db = initialize.GormMysql() // gorm连接数据库
	if global.QY_Db != nil {
		//initialize.RegisterTables() // 初始化表
		//err := menu.InitRouterDb()
		//if err != nil {
		//	return
		//}
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
