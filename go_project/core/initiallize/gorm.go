package initiallize

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"louis/global"
	"os"
	"time"
)

var Gorm = new(_gorm)

type _gorm struct{}

func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(logger.Writer(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DBBASE
	switch global.LOUIS_CONFIG.System.DbType {
	case "mysql":
		logMode = &global.LOUIS_CONFIG.Mysql

	default:
		logMode = &global.LOUIS_CONFIG.Mysql
	}
	switch logMode.GetLogMode() {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
