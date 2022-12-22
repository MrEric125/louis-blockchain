package initiallize

import (
	"encoding/json"
	"fmt"
	"louis/cache"
	"louis/global"
	"louis/utils"
)

func OuterInit() {

	dr, err := utils.ParseDuration(global.LOUIS_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.LOUIS_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = cache.NewCache(
		cache.SetDefaultExpire(dr),
	)
	jsonU, _ := json.Marshal(&global.LOUIS_CONFIG)
	fmt.Println(string(jsonU))

}
