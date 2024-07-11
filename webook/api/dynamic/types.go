package dynamic

import (
	"github.com/gin-gonic/gin"
	v1 "webook/api/dynamic/v1"
	"webook/pkg/ginx"
)

type IDynamicV1 interface {
	//发布文章视频
	RegisterRoutes(server *gin.Engine)
	PublishDynamicV1(ctx *gin.Context, req v1.PublishDynamicReq) (ginx.Result, error)
}
