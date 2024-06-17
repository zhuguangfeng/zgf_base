package repository

import (
	"context"
	"zyz_jike/search/domain"
)

type ArticleRepository interface {
	InputArticle(ctx context.Context, msg domain.Article) error
	SearchArticle(ctx context.Context, keywords []string) ([]domain.Article, error)
}

type UserRepository interface {
	InputUser(ctx context.Context, msg domain.User) error
	SearchUser(ctx context.Context, keywords []string) ([]domain.User, error)
}
