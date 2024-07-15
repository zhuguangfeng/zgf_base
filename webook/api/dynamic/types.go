package dynamic

import (
	"github.com/gin-gonic/gin"
	v1 "webook/api/dynamic/v1"
	"webook/pkg/ginx"
)

type IDynamicV1 interface {
	//发布文章视频
	RegisterRoutes(server *gin.Engine)
	PublishDynamic(ctx *gin.Context, req v1.PublishDynamicReq) (ginx.Result, error)
	SearchDynamicList(ctx *gin.Context, req v1.SearchDynamicListReq) (ginx.Result, error)
}
