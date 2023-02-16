package main

import (
	"fmt"
	"go-generate/core"
	"go-generate/global"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	fmt.Println(global.GVA_VP)
}
