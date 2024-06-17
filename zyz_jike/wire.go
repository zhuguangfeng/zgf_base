//go:build wireinject

package main

import (
	"github.com/google/wire"
	"zyz_jike/internal/repository"
	"zyz_jike/internal/repository/dao"
	"zyz_jike/internal/service"
	"zyz_jike/internal/web/article"
	ijwt "zyz_jike/internal/web/jwt"
	"zyz_jike/internal/web/user"
	"zyz_jike/ioc"
)

func InitWebServer() *App {
	wire.Build(
		//第三方依赖
		ioc.InitLogger,
		ioc.InitDB,
		ioc.InitRedis,
		ioc.InitGinMiddleware,
		ioc.InitWebServer,

		dao.NewUserDao,
		dao.NewArticleDao,

		repository.NewUserRepository,
		repository.NewArticleRepository,

		service.NewUserService,
		service.NewArticleService,

		ijwt.NewRedisJwtHandler,
		user.NewUserHandler,
		article.NewArticleHandler,

		wire.Struct(new(App), "*"),
	)
	return new(App)
}
