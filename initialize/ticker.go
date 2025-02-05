package initialize

import (
	"file_service/global"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Ticker() {
	ticker := time.NewTicker(24 * time.Hour) // 每 24 小时执行一次
	defer ticker.Stop()
	// 立即执行一次备份
	go backupDatabase()
	// 进入定时循环
	for {
		select {
		case <-ticker.C:
			go backupDatabase()
		}
	}
}

func backupDatabase() {
	timestamp := time.Now().Format("20060102_150405")
	backupFile := fmt.Sprintf("backup_mysql_%s.sql", timestamp)
	cmd := exec.Command("mysqldump", "-u", "root", "-p"+global.QY_CONFIG.Mysql.Password, global.QY_CONFIG.Mysql.Dbname)
	output, err := cmd.Output()
	if err != nil {
		global.QY_LOG.Error("数据库备份失败" + err.Error())
		return
	}
	err = os.WriteFile(backupFile, output, 0644)
	if err != nil {
		global.QY_LOG.Error("备份文件保存失败" + err.Error())
		return
	}
	global.QY_LOG.Info("数据备份成功")
}
