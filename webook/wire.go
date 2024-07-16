//go:build wireinject

package main

import (
	"github.com/google/wire"
	dynamic2 "webook/internal/events/dynamic"
	"webook/internal/repository"
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
	"webook/internal/service"
	"webook/internal/web/dynamic"
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
		dynamic.NewDynamicHandlerV1,

		dynamic2.NewDynamicConsumer,
		ioc.NewConsumers,
		ioc.InitSaramaSyncProducer,
		dynamic2.NewSaramaProducer,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
