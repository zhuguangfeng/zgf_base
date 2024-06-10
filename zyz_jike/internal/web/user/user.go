package user

import (
	"github.com/gin-gonic/gin"
	"zyz_jike/pkg/ginx"
)

type UserHandler interface {
	Register(ctx *gin.Context) (ginx.Result, error)
}

type userHandler struct {
}

func (u *userHandler) RegisterRoutes(server *gin.Engine) {
	appGroup := server.Group("/app/user")
	{
		appGroup.POST("register")
	}
}

func (u *userHandler) Register(ctx *gin.Context, req UserRegisterReq) (ginx.Result, error) {
	return ginx.Result{}, nil
}
