package main

import (
	"louis/core"
	"louis/core/initiallize"
	"louis/global"
)

func main() {
	// 初始化viper
	global.LOUIS_VP = core.Viper()
	initiallize.OuterInit()

}
