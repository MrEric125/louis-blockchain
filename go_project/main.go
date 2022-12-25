package main

import (
	"louis/core"
	"louis/core/initiallize"
)

func main() {
	// 初始化viper
	core.Viper()
	// 初始化zap
	initiallize.ZapInit()
	initiallize.OuterInit()
	// 初始化数据库
	initiallize.SqlInit()
	// 初始化路由
	initiallize.DoInitRouters()

}
