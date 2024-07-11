package dynamic

import (
	"github.com/gin-gonic/gin"
	v1 "webook/api/dynamic/v1"
	"webook/internal/domain"
	"webook/internal/errs"
	"webook/pkg/ginx"
	"webook/pkg/logger"
)

func (c *DynamicControllerV1) PublishDynamicV1(ctx *gin.Context, req v1.PublishDynamicReq) (ginx.Result, error) {
	err := c.svc.PublishDynamic(ctx, domain.Dynamic{
		Title:     req.Title,
		Content:   req.Content,
		Resources: req.Resources,
	})
	if err != nil {
		c.l.Error("发布资源失败", logger.Field{Key: "req", Val: req}, logger.Error(err))
		return ginx.Result{
			Code:    errs.ArticleInternalServerError,
			Message: "系统错误",
		}, nil
	}
	return ginx.Result{
		Code:    errs.Success,
		Message: "ok",
	}, err
}
