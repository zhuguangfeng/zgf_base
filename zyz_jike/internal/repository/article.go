package repository

import (
	"context"
	"zyz_jike/internal/domain"
	"zyz_jike/internal/repository/dao"
)

type ArticleRepository interface {
	CreateArticle(ctx context.Context, article domain.Article) (domain.Article, error)
	UpdateArticle(ctx context.Context, article domain.Article) error
	DeleteArticle(ctx context.Context, id []int64) error
	GetArticle(ctx context.Context, id int64) (domain.Article, error)
}
type articleRepository struct {
	dao dao.ArticleDao
}

func NewArticleRepository(dao dao.ArticleDao) ArticleRepository {
	return &articleRepository{
		dao: dao,
	}
}

func (repo *articleRepository) DeleteArticle(ctx context.Context, ids []int64) error {
	return repo.dao.DeleteArticle(ctx, ids)
}

func (repo *articleRepository) CreateArticle(ctx context.Context, article domain.Article) (domain.Article, error) {
	art, err := repo.dao.InsertArticle(ctx, repo.toEntity(article))
	if err != nil {
		return domain.Article{}, err
	}
	return repo.toDomain(art), nil
}

func (repo *articleRepository) UpdateArticle(ctx context.Context, article domain.Article) error {
	return repo.dao.UpdateArticle(ctx, repo.toEntity(article))
}

func (repo *articleRepository) GetArticle(ctx context.Context, id int64) (domain.Article, error) {
	daoArt, err := repo.dao.FindArticleById(ctx, id)
	if err != nil {
		return domain.Article{}, err
	}
	return repo.toDomain(daoArt), nil
}

func (repo *articleRepository) toEntity(art domain.Article) dao.Article {
	return dao.Article{
		Model: &dao.Model{
			ID: art.Id,
		},
		Uid:             art.Author.Id,
		Category:        art.Category,
		ArticleCategory: art.ArticleCategory,
		Title:           art.Title,
		Content:         art.Content,
		RichText:        art.RichText,
		Pic:             art.Pic,
		Pics:            art.Pics,
		Status:          art.Status,
	}
}

func (repo *articleRepository) toDomain(art dao.Article) domain.Article {
	return domain.Article{
		Id: art.ID,
		Author: domain.Author{
			Id: art.Uid,
		},
		Category:        art.Category,
		ArticleCategory: art.ArticleCategory,
		Title:           art.Title,
		Content:         art.Content,
		RichText:        art.RichText,
		Pic:             art.Pic,
		Pics:            art.Pics,
		Status:          art.Status,
		CreatedAt:       art.CreatedAt,
	}
}
