package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"zyz_jike/pkg/grpcx"
	"zyz_jike/search/events"
)

func main() {
	initViperWatch()
	app := Init()
	for _, c := range app.consumers {
		err := c.Start()
		if err != nil {
			panic(err)
		}
	}
	err := app.server.Serve()
	if err != nil {
		panic(err)
	}
}
func initViperWatch() {
	cfile := pflag.String("config", "config/dev.yaml", "配置文件路径")
	pflag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(*cfile)
	viper.WatchConfig()

	//读取配置
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

type App struct {
	server    *grpcx.Server
	consumers []events.Consumer
}
