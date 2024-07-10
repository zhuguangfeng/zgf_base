package dynamic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	v1 "webook/api/dynamic/v1"
	"webook/internal/domain"
	"webook/pkg/ginx"
)

func (c *DynamicControllerV1) PublishDynamicV1(ctx *gin.Context, req v1.PublishDynamicReq) (ginx.Result, error) {
	err := c.svc.PublishDynamic(ctx, domain.Dynamic{
		Title:     req.Title,
		Content:   req.Content,
		Resources: req.Resources,
	})
	fmt.Println(err)
}
