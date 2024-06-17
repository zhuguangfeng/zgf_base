package article

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"zyz_jike/internal/codes"
	"zyz_jike/internal/domain"
	"zyz_jike/internal/service"
	"zyz_jike/internal/web/jwt"
	"zyz_jike/pkg/ginx"
	"zyz_jike/pkg/logger"
)

type ArticleHandler struct {
	artSvc service.ArticleService

	l logger.Logger
}

func NewArticleHandler(artSvc service.ArticleService, l logger.Logger) *ArticleHandler {
	return &ArticleHandler{
		artSvc: artSvc,
		l:      l,
	}
}
func (hdl *ArticleHandler) RegisterRoutes(server *gin.Engine) {
	appGroup := server.Group("/app/article")
	{
		appGroup.POST("/publish", ginx.WrapBodyAndClaims(hdl.PublishArticle))
		appGroup.GET("/detail", hdl.ArticleDetail)
	}
}

func (hdl *ArticleHandler) PublishArticle(ctx *gin.Context, req PublishArticleReq, uc jwt.UserClaims) (ginx.Result, error) {

	err := hdl.artSvc.CreateArticle(ctx, domain.Article{
		Title:   req.Title,
		Content: req.Content,
		Pic:     req.Pics[0],
		Pics:    req.Pics,
		Author: domain.Author{
			Id: uc.Uid,
		},
		Status: domain.ArticlePublished,
	})
	if err != nil {
		return ginx.Result{
			Code:    codes.ArticleInternalServerError,
			Message: "系统错误",
		}, err
	}
	return ginx.Result{
		Code:    codes.Success,
		Message: "success",
	}, nil
}

func (hdl *ArticleHandler) ArticleDetail(ctx *gin.Context) {
	id := ctx.Query("id")
	idi, _ := strconv.Atoi(id)
	art, err := hdl.artSvc.ArticleDetail(ctx, int64(idi))
	if err != nil {
		hdl.l.Error("查询文章详情失败", logger.Error(err))
		ctx.JSON(http.StatusOK, ginx.Result{
			Code:    codes.ArticleInternalServerError,
			Message: "系统错误",
		})
		return
	}
	vo := ArticleVo{
		Id:       art.Id,
		Title:    art.Title,
		Content:  art.Content,
		Pic:      art.Pic,
		Pics:     art.Pics,
		Uid:      art.Author.Id,
		Nickname: art.Author.Nickname,
	}
	ctx.JSON(http.StatusOK, ginx.Result{
		Code:    codes.Success,
		Message: "success",
		Data:    vo,
	})
	return
}
