package main

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	initViperWatch()

	//l := ioc.InitLogger()
	//db := ioc.InitDB()
	//dao := dao2.NewUserDao(db)
	//redisCmd := ioc.InitRedis()
	//userRepo := repository.NewUserRepository(dao)
	//usersvc := service.NewUserService(userRepo)
	//hdl := ijwt.NewRedisJwtHandler(redisCmd)
	//userHdl := user.NewUserHandler(hdl, usersvc, l)
	//
	//mdls := ioc.InitGinMiddleware(redisCmd, hdl, l)
	//r := ioc.InitWebServer(mdls, userHdl)
	app := InitWebServer()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	err := app.Server.Run(":8888") // 监听并在 0.0.0.0:8080 上启动服务
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
