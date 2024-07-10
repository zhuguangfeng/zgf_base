package article

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	searchv1 "zyz_jike/api/proto/gen/search/v1"
	"zyz_jike/internal/codes"
	"zyz_jike/internal/domain"
	"zyz_jike/internal/event/article"
	"zyz_jike/internal/service"
	"zyz_jike/internal/web/jwt"
	"zyz_jike/pkg/ginx"
	"zyz_jike/pkg/logger"
)

type ArticleHandler struct {
	artSvc       service.ArticleService
	search       searchv1.SearchServiceClient
	sync         searchv1.SyncServiceClient
	syncArtEvent article.Producer
	l            logger.Logger
}

func NewArticleHandler(artSvc service.ArticleService, syncArtEvent article.Producer, search searchv1.SearchServiceClient, sync searchv1.SyncServiceClient, l logger.Logger) *ArticleHandler {
	return &ArticleHandler{
		artSvc:       artSvc,
		search:       search,
		syncArtEvent: syncArtEvent,
		sync:         sync,
		l:            l,
	}
}
func (hdl *ArticleHandler) RegisterRoutes(server *gin.Engine) {
	appGroup := server.Group("/app/dynamic")
	{
		appGroup.POST("/publish", ginx.WrapBodyAndClaims(hdl.PublishArticle))
		appGroup.GET("/detail", hdl.ArticleDetail)
		appGroup.POST("/search", ginx.WrapBodyAndClaims(hdl.ArticleSearch))
	}
}

func (hdl *ArticleHandler) PublishArticle(ctx *gin.Context, req PublishArticleReq, uc jwt.UserClaims) (ginx.Result, error) {
	art, err := hdl.artSvc.CreateArticle(ctx, domain.Article{
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

	//异步同步文章到es
	err = hdl.syncArtEvent.ProduceSyncArticleEvent(article.ArticleEvent{
		Id:              art.Id,
		Title:           req.Title,
		Content:         req.Content,
		Status:          art.Status,
		Pic:             art.Pic,
		Pics:            art.Pics,
		Category:        art.Category,
		ArticleCategory: art.ArticleCategory,
		RichText:        art.RichText,
		CreatedAt:       art.CreatedAt.Format("2006-01-02 15:04:05"),
		Uid:             art.Author.Id,
	})
	if err != nil {
		hdl.l.Error("同步失败", logger.Error(err))
	}

	//同步 同步文章到es
	//searchV1rt := &searchv1.Article{
	//	Id:              art.Id,
	//	Title:           req.Title,
	//	Content:         req.Content,
	//	Status:          int32(art.Status),
	//	Pic:             art.Pic,
	//	Pics:            art.Pics,
	//	Category:        int32(art.Category),
	//	ArticleCategory: int32(art.ArticleCategory),
	//	RichText:        art.RichText,
	//	CreatedAt:       art.CreatedAt.Format("2006-01-02 15:04:05"),
	//	Uid:             art.Author.Id,
	//}
	//_, err = hdl.sync.InputArticle(ctx, &searchv1.InputArticleRequest{
	//	Article: searchV1rt,
	//})
	//if err != nil {
	//	hdl.l.Error("同步失败")
	//}
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

func (hdl *ArticleHandler) ArticleSearch(ctx *gin.Context, req SearchArticleReq, uc jwt.UserClaims) (ginx.Result, error) {
	resp, err := hdl.search.Search(ctx, &searchv1.SearchRequest{
		Expression: req.Expression,
	})
	fmt.Println(resp)
	if err != nil {
		hdl.l.Error("搜索文章失败", logger.String("搜索条件", req.Expression), logger.Error(err))
		return ginx.Result{
			Code:    codes.ArticleInternalServerError,
			Message: "系统错误",
		}, err
	}
	return ginx.Result{
		Code:    codes.Success,
		Message: "success",
		Data:    resp,
	}, nil
}
