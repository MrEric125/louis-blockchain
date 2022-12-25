package initiallize

import (
	"encoding/json"
	"louis/cache"
	"louis/global"
	"louis/utils"
)

type Outer struct {
}

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

	logger.Info(string(jsonU))
	logger.Info("==========log===============")

}
