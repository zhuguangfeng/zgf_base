//go:build wireinject

package main

import (
	"github.com/google/wire"
	"webook/internal/controller/dynamic"
	"webook/internal/events"
	"webook/internal/repository"
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/ioc"
)

func InitWebService() *App {
	wire.Build(
		ioc.InitLogger,
		ioc.InitDB,
		ioc.InitRedis,
		ioc.InitEsClient,
		ioc.InitKafka,
		ioc.InitGinMiddleware,
		ioc.InitWebServer,

		dao.NewDynamicDao,
		dao.NewOlivereDynamicEsDao,
		cache.NewDynamicCache,
		repository.NewDynamicRepository,
		service.NewDynamicService,
		dynamic.NewDynamicControllerV1,

		events.NewDynamicConsumer,
		ioc.NewConsumers,
		ioc.InitSaramaSyncProducer,
		events.NewSaramaProducer,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
