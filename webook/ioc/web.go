package ioc

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
	"webook/api/dynamic"
)

func InitWebServer(mdls []gin.HandlerFunc, dynamicHdl dynamic.IDynamicV1) *gin.Engine {
	server := gin.Default()
	server.Use(mdls...)

	dynamicHdl.RegisterRoutes(server)
	return server
}

func InitGinMiddleware(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		//跨域配置
		cors.New(cors.Config{
			//AllowOrigins: []string{"http://localhost:3000"}, //枚举允许那些跨域请求
			AllowHeaders:  []string{"Content-Type", "Authorization"}, //允许的请求头
			ExposeHeaders: []string{"x-jwt-token"},                   //允许前端访问你的后端响应中带的头部
			AllowOriginFunc: func(origin string) bool { //请求地址如果包含localhost可以请求
				return strings.Contains(origin, "localhost")
			},
			MaxAge: time.Hour * 12,
		}),
		//middleware.NewLoginJwtMiddlewareBuilder(hdl).CheckLogin(),
	}
}
