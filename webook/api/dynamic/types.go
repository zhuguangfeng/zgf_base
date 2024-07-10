package dynamic

import (
	"github.com/gin-gonic/gin"
	v1 "webook/api/dynamic/v1"
	"webook/pkg/ginx"
)

type IDynamicV1 interface {
	PublishDynamicV1(ctx *gin.Context, req v1.PublishDynamicReq) (ginx.Result, error)
}
