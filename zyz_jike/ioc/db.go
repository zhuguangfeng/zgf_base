package ioc

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zyz_jike/internal/repository/dao"
	"zyz_jike/pkg/logger"
)

func InitDB() *gorm.DB {
	type Config struct {
		DSN string `yaml:"dsn"`
	}
	var cfg = Config{
		DSN: "root:root@tcp(localhost:3316)/zgf?charset=utf8&parseTime=true",
	}
	err := viper.UnmarshalKey("db", &cfg)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(cfg.DSN), &gorm.Config{
		//Logger: glogger.New(goormLoggerFunc(l.Debug), glogger.Config{
		//	// 慢查询
		//	SlowThreshold: 0,
		//	LogLevel:      glogger.Info,
		//}),
	})
	if err != nil {
		panic(err)
	}
	if err = dao.InitTables(db); err != nil {
		panic(err)
	}
	return db
}

type goormLoggerFunc func(msg string, fields ...logger.Field)

func (g goormLoggerFunc) Printf(s string, i ...interface{}) {
	g(s, logger.Field{Key: "args", Val: i})
}
