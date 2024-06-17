package repository

import (
	"context"
	"github.com/ecodeclub/ekit/slice"
	"zyz_jike/search/domain"
	"zyz_jike/search/repository/dao"
)

type articleRepository struct {
	dao dao.ArticleDao
}

func (repo *articleRepository) InputArticle(ctx context.Context, msg domain.Article) error {
	return repo.dao.InputArticle(ctx, dao.Article{
		Id:              msg.Id,
		Uid:             msg.Uid,
		Category:        msg.Category,
		ArticleCategory: msg.ArticleCategory,
		Title:           msg.Title,
		Content:         msg.Content,
		RichText:        msg.RichText,
		Pic:             msg.Pic,
		Pics:            msg.Pics,
		Status:          msg.Status,
	})
}

func (repo *articleRepository) SearchArticle(ctx context.Context, keywords []string) ([]domain.Article, error) {
	arts, err := repo.dao.Search(ctx, []int64{}, keywords)
	if err != nil {
		return nil, err
	}
	return slice.Map(arts, func(idx int, src dao.Article) domain.Article {
		return domain.Article{
			Id:              src.Id,
			Uid:             src.Uid,
			Category:        src.Category,
			ArticleCategory: src.ArticleCategory,
			Title:           src.Title,
			Content:         src.Content,
			RichText:        src.RichText,
			Pic:             src.Pic,
			Pics:            src.Pics,
			Status:          src.Status,
		}
	}), nil
}

func NewArticleRepository(d dao.ArticleDao) ArticleRepository {
	return &articleRepository{
		dao: d,
	}
}
