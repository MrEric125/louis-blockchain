package global

import (
	"louis/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// gorm https://gorm.io/
// redis https://github.com/go-redis/redis
// viper https://github.com/spf13/viper
// zap https://github.com/uber-go/zap
var (
	LOUIS_DB            *gorm.DB
	LOUIS_DBList        map[string]*gorm.DB
	LOUIS_REDIS         *redis.Client
	LOUIS_CONFIG        config.Server
	LOUIS_VP            *viper.Viper
	LOUIS_LOG           *zap.Logger
	DEFAULT_CONFIG_TYPE string = "yaml"
)
