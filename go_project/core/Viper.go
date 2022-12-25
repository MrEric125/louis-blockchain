package core

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"louis/core/internal"
	"louis/global"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/*
*
读取yaml 文件工具类初始化
*/
func Viper(path ...string) *viper.Viper {

	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()

		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
				case gin.TestMode:
					config = internal.ConfigTestFile
				}
			} else {
				config = configEnv
			}
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)

	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType(global.DEFAULT_CONFIG_TYPE)
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err = v.Unmarshal(&global.LOUIS_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	// 给global.LOUIS_CONFIG 配置项赋值
	if err = v.Unmarshal(&global.LOUIS_CONFIG); err != nil {
		fmt.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.LOUIS_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v

}
