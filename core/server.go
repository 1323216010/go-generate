package core

import (
	"fmt"
	"time"

	"go-generate/global"
	"go-generate/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	// 从db加载jwt数据
	if global.YAN_DB != nil {
		//system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.YAN_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.YAN_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	swagger地址:http://127.0.0.1%s/swagger/index.html
`, address)
	global.YAN_LOG.Error(s.ListenAndServe().Error())
}
