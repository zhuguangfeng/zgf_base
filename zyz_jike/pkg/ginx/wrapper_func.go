package ginx

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"zyz_jike/pkg/logger"
)

var L = logger.NewNopLogger()

// 请求体处理方法
func WrapBody[Req any](bizFn func(ctx *gin.Context, req Req) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			L.Error("输入错误", logger.Error(err))
			return
		}
		L.Debug("输出参数", logger.Field{Key: "req", Val: req})
		res, err := bizFn(ctx, req)
		if err != nil {
			L.Error("执行业务逻辑失败", logger.Error(err))

		}
		ctx.JSON(http.StatusOK, res)
	}

}

func WrapBodyAndClaims[Req any, Claims jwt.Claims](bizFn func(ctx *gin.Context, req Req, uc Claims) (Result, error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Req
		if err := ctx.Bind(&req); err != nil {
			L.Error("输入错误", logger.Error(err))
			return
		}
		L.Debug("输入参数", logger.Field{Key: "req", Val: req})

		val, ok := ctx.Get("user")
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		uc, ok := val.(Claims)
		if !ok {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		res, err := bizFn(ctx, req, uc)
		if err != nil {
			L.Error("执行业务逻辑失败", logger.Error(err))

		}
		ctx.JSON(http.StatusOK, res)
	}

}
