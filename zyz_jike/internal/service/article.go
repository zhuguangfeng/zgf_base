package service

import (
	"context"
	"zyz_jike/internal/domain"
	"zyz_jike/internal/repository"
	"zyz_jike/pkg/logger"
)

type ArticleService interface {
	CreateArticle(ctx context.Context, art domain.Article) error
	DeleteArticle(ctx context.Context, ids []int64) error
	ArticleDetail(ctx context.Context, id int64) (domain.Article, error)
}
type articleService struct {
	artRepo repository.ArticleRepository
	userReo repository.UserRepository
	l       logger.Logger
}

func NewArticleService(artRepo repository.ArticleRepository, userReo repository.UserRepository, l logger.Logger) ArticleService {
	return &articleService{
		artRepo: artRepo,
		userReo: userReo,
		l:       l,
	}
}
func (svc *articleService) CreateArticle(ctx context.Context, art domain.Article) error {
	return svc.artRepo.CreateArticle(ctx, art)
}

func (svc *articleService) DeleteArticle(ctx context.Context, ids []int64) error {
	return svc.artRepo.DeleteArticle(ctx, ids)
}

func (svc *articleService) ArticleDetail(ctx context.Context, id int64) (domain.Article, error) {
	art, err := svc.artRepo.GetArticle(ctx, id)
	if err != nil {
		return domain.Article{}, err
	}
	daoU, err := svc.userReo.FindUserById(ctx, art.Id)
	if err != nil {
		svc.l.Error("获取文章详情 查询文章作者失败", logger.Int64("art_id", id), logger.Error(err))
		return art, nil
	}
	art.Author.Nickname = daoU.Nickname
	return art, nil
}
