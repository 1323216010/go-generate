package initialize

import (
	"fmt"

	"go-generate/utils/plugin"

	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)
	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	/*	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())*/
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	/*	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.YAN_CONFIG.Email.To,
		global.YAN_CONFIG.Email.From,
		global.YAN_CONFIG.Email.Host,
		global.YAN_CONFIG.Email.Secret,
		global.YAN_CONFIG.Email.Nickname,
		global.YAN_CONFIG.Email.Port,
		global.YAN_CONFIG.Email.IsSSL,
	))*/
}
