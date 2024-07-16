package dynamic

import (
	"github.com/gin-gonic/gin"
	"webook/api/dynamic"
	v1 "webook/api/dynamic/v1"
	"webook/internal/service"
	"webook/pkg/ginx"
	"webook/pkg/logger"
)

type DynamicHandlerV1 struct {
	svc service.DynamicService
	l   logger.Logger
}

func NewDynamicHandlerV1(svc service.DynamicService, l logger.Logger) dynamic.IDynamicV1 {
	return &DynamicHandlerV1{
		svc: svc,
		l:   l,
	}
}

func (c *DynamicHandlerV1) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/v1/dynamic")
	{
		g.POST("/publish", ginx.WrapBody[v1.PublishDynamicReq](c.PublishDynamic))
		g.POST("/search", ginx.WrapBody[v1.SearchDynamicListReq](c.SearchDynamicList))
	}

}
