package ioc

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"zyz_jike/pkg/logger"
)

func InitLogger() logger.Logger {
	cfg := zap.NewDevelopmentConfig()
	err := viper.UnmarshalKey("log", &cfg)
	if err != nil {
		panic(err)
	}
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger.NewZapLogger(l)
}
