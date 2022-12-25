package main

import (
	"louis/core"
	"louis/core/initiallize"
)

func main() {
	// 初始化viper
	core.Viper()

	initiallize.ZapInit()
	initiallize.OuterInit()
	initiallize.SqlInit()

	rout := initiallize.Routers{}

	rout.DoInitRouters()

}
