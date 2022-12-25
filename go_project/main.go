package main

import (
	"louis/core"
	"louis/core/initiallize"
	"louis/global"
)

func main() {
	// 初始化viper
	global.LOUIS_VP = core.Viper()

	initiallize.ZapInit()
	initiallize.OuterInit()
	initiallize.SqlInit()

	rout := initiallize.Routers{}

	rout.DoInitRouters()

}
