package main

import (
	"louis/core"
	"louis/global"
)

func main() {
	// 初始化viper
	global.LOUIS_VP = core.Viper()

}
