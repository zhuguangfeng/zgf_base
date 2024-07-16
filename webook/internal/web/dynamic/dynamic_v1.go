package dynamic

import (
	"github.com/gin-gonic/gin"
	v1 "webook/api/dynamic/v1"
	"webook/internal/domain"
	"webook/internal/errs"
	"webook/pkg/ginx"
	"webook/pkg/logger"
)

func (c *DynamicHandlerV1) PublishDynamic(ctx *gin.Context, req v1.PublishDynamicReq) (ginx.Result, error) {
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
		}, err
	}

	return ginx.Result{
		Code:    errs.Success,
		Message: "ok",
	}, nil
}

func (c *DynamicHandlerV1) SearchDynamicList(ctx *gin.Context, req v1.SearchDynamicListReq) (ginx.Result, error) {
	data, err := c.svc.SearchDynamicPage(ctx, req)
	if err != nil {
		return ginx.Result{
			Code:    errs.ArticleInternalServerError,
			Message: "系统错误",
		}, err
	}

	res := make([]v1.Dynamic, 0, len(data))
	if len(data) > 0 {
		for _, dnc := range data {
			res = append(res, c.toVo(dnc))
		}
	}

	return ginx.Result{
		Code:    errs.Success,
		Message: "ok",
		Data:    res,
	}, nil
}

func (c *DynamicHandlerV1) toVo(dnc domain.Dynamic) v1.Dynamic {
	return v1.Dynamic{
		Id:         dnc.Id,
		Title:      dnc.Title,
		Content:    dnc.Content,
		Category:   dnc.Category,
		Resources:  dnc.Resources,
		CreateTime: dnc.CreateTime,
		UpdateTime: dnc.CreateTime,
	}
}
