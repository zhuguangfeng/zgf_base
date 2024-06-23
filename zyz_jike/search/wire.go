//go:build wireinject

package main

import (
	"github.com/google/wire"
	"zyz_jike/search/events"

	"zyz_jike/search/grpc"
	"zyz_jike/search/ioc"
	"zyz_jike/search/repository"
	"zyz_jike/search/repository/dao"
	"zyz_jike/search/service"
)

var serviceProviderSet = wire.NewSet(
	dao.NewUserElasticDao,
	dao.NewArticleElasticDao,
	repository.NewUserRepository,
	repository.NewArticleRepository,
	service.NewSyncService,
	service.NewSearchService,
)

var thirdProvider = wire.NewSet(
	ioc.InitEsClient,
	ioc.InitEtcdClient,
	ioc.InitLogger,
	ioc.InitKafka,
)

func Init() *App {
	wire.Build(
		thirdProvider,
		serviceProviderSet,
		grpc.NewSyncServiceServer,
		grpc.NewSearchService,
		events.NewArticleConsumer,
		ioc.InitGrpcxServer,
		ioc.NewConsumers,
		wire.Struct(new(App), "*"),
	)
	return new(App)
}
