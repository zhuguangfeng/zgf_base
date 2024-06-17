package user

import (
	"fmt"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-gonic/gin"
	"net/http"
	"zyz_jike/internal/codes"
	"zyz_jike/internal/domain"
	"zyz_jike/internal/errs"
	"zyz_jike/internal/service"
	ijwt "zyz_jike/internal/web/jwt"
	"zyz_jike/pkg/ginx"
	"zyz_jike/pkg/logger"
)

const phoneRegexPattern = `^1[3-9]\d{9}$`

//type UserHandler interface {
//	Register(ctx *gin.Context, req UserRegisterReq) (ginx.Result, error)
//}

type UserHandler struct {
	ijwt.Handler
	phoneRegex *regexp.Regexp
	userSvc    service.UserService
	l          logger.Logger
}

func NewUserHandler(hdl ijwt.Handler, userSvc service.UserService, l logger.Logger) *UserHandler {
	return &UserHandler{
		Handler:    hdl,
		phoneRegex: regexp.MustCompile(phoneRegexPattern, regexp.None),
		userSvc:    userSvc,
		l:          l,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	appGroup := server.Group("/app/user")
	{
		appGroup.POST("/register", ginx.WrapBody[UserRegisterReq](h.Register))
		appGroup.POST("/pwd/login", ginx.WrapBody[PwdLoginReq](h.PwdLogin))
		appGroup.GET("/logout", h.Logout)
	}
}

// 用户注册
func (h *UserHandler) Register(ctx *gin.Context, req UserRegisterReq) (ginx.Result, error) {
	if req.Password != req.RePassword {
		return ginx.Result{
			Code:    codes.UserInvalidInput,
			Message: "两次输入密码不一致",
		}, nil
	}
	isPhone, err := h.phoneRegex.MatchString(req.Phone)
	if err != nil {
		return ginx.Result{
			Code:    codes.UserInternalServerError,
			Message: "系统错误",
		}, err
	}
	if !isPhone {
		return ginx.Result{
			Code:    codes.UserInvalidInput,
			Message: "非法的手机号码",
		}, nil
	}
	err = h.userSvc.UserRegister(ctx, domain.User{
		Phone:    req.Phone,
		Password: req.Password,
	})
	switch err {
	case nil:
		return ginx.Result{
			Code:    codes.Success,
			Message: "success",
		}, nil
	case errs.ErrDuplicatePhone:
		return ginx.Result{
			Code:    codes.UserDuplicatePhone,
			Message: "手机号码已存在",
		}, err
	default:
		return ginx.Result{
			Code:    codes.UserInternalServerError,
			Message: "系统错误",
		}, err

	}
}

// 手机密码登录
func (h *UserHandler) PwdLogin(ctx *gin.Context, req PwdLoginReq) (ginx.Result, error) {
	ifPhone, err := h.phoneRegex.MatchString(req.Phone)
	if err != nil {
		return ginx.Result{
			Code:    codes.UserInternalServerError,
			Message: "系统错误",
		}, err
	}
	if !ifPhone {
		return ginx.Result{
			Code:    codes.UserInvalidInput,
			Message: "非法的手机号码",
		}, nil
	}

	user, err := h.userSvc.PwdLogin(ctx, req.Phone, req.Password)
	switch err {
	case nil:
		err = h.SetLoginToken(ctx, user.Id)
		if err != nil {
			return ginx.Result{
				Code:    codes.UserInternalServerError,
				Message: "系统错误",
			}, err
		}
		return ginx.Result{
			Code:    codes.Success,
			Message: "success",
		}, nil
	case errs.ErrUserNotFound:
		return ginx.Result{
			Code:    codes.UserInvalidInput,
			Message: "手机号码未注册",
		}, nil
	case errs.ErrInvalidPassword:
		return ginx.Result{
			Code:    codes.UserInvalidInput,
			Message: "密码错误",
		}, nil
	default:
		return ginx.Result{
			Code:    codes.UserInternalServerError,
			Message: "系统错误",
		}, err
	}
}

func (h *UserHandler) Logout(ctx *gin.Context) {
	fmt.Println(ctx.GetHeader("Authorization"))
	err := h.ClearToken(ctx)
	if err != nil {
		uc := ctx.MustGet("user").(ijwt.UserClaims)
		h.l.Error("退出登录失败", logger.Int64("uid", uc.Uid), logger.Error(err))
		ctx.JSON(http.StatusOK, ginx.Result{
			Code:    codes.UserInternalServerError,
			Message: "系统错误",
		})
		return
	}
	ctx.JSON(http.StatusOK, ginx.Result{
		Code:    codes.Success,
		Message: "退出登录成功",
	})
}
