package main

import (
	"go-generate/core"
	"go-generate/global"
	"go-generate/initialize"

	"go.uber.org/zap"
)

func main() {
	global.YAN_VP = core.Viper() // 初始化Viper
	global.YAN_LOG = core.Zap()  // 初始化zap日志库
	zap.ReplaceGlobals(global.YAN_LOG)
	global.YAN_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.YAN_DB != nil {
		initialize.RegisterTables(global.YAN_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.YAN_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
