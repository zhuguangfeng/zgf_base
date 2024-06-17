package dao

import "context"

type UserDao interface {
	InputUser(ctx context.Context, user User) error
	Search(ctx context.Context, keywords []string) ([]User, error)
}

type ArticleDao interface {
	InputArticle(ctx context.Context, article Article) error
	Search(ctx context.Context, tagArtIds []int64, keywords []string) ([]Article, error)
}

type TagDao interface {
	Search(ctx context.Context, uid int64, biz string, keywords []string) ([]int64, error)
}
