package initiallize

import (
	"encoding/json"
	"fmt"
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

	var jsonObj map[string]interface{}
	err = json.Unmarshal([]byte(jsonU), &jsonObj)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	formattedJson, err := json.MarshalIndent(jsonObj, "", "    ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	logger.Info(string(formattedJson))
	logger.Info("==========log===============")

}
