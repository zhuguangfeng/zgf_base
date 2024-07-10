package dynamic

import (
	"github.com/gin-gonic/gin"
	"webook/api/dynamic"
	v1 "webook/api/dynamic/v1"
	"webook/internal/service"
	"webook/pkg/ginx"
)

type DynamicControllerV1 struct {
	svc service.DynamicService
}

func NewDynamicControllerV1(svc service.DynamicService) dynamic.IDynamicV1 {
	return &DynamicControllerV1{
		svc: svc,
	}
}

func (c *DynamicControllerV1) RegisterRoutes(server *gin.Engine) {
	g := server.Group("/v1/dynamic")
	{
		g.POST("/publish", ginx.WrapBody[v1.PublishDynamicReq](c.PublishDynamicV1))
	}

}
