package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"louis/core"
	"louis/core/initiallize"
	"louis/global"
)

func main() {
	// 初始化viper
	global.LOUIS_VP = core.Viper()

	initiallize.OuterInit()

	rout := initiallize.Routers{}

	rout.DoRouters()

	switch global.LOUIS_CONFIG.System.DbType {
	case "mysql":
		GormMysql()
	}
}

func GormMysql() *gorm.DB {
	m := global.LOUIS_CONFIG.Mysql
	if m.Dbname == "" {
		global.LOUIS_LOG.Info("数据库连接未找到")
		return nil
	}

	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(),
		DefaultStringSize:         191,
		SkipInitializeWithVersion: false,
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), initiallize.Gorm.Config(m.Prefix, m.Singular))
	if err != nil {
		panic(err)
		return nil
	}
	db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(m.MaxIdleConns)
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	return db

}
