package service

import (
	"context"
	"zyz_jike/search/domain"
	"zyz_jike/search/repository"
)

type SyncService interface {
	InputArticle(ctx context.Context, article domain.Article) error
	InputUser(ctx context.Context, user domain.User) error
}

type syncService struct {
	userRepo repository.UserRepository
	artRepo  repository.ArticleRepository
}

func (s *syncService) InputArticle(ctx context.Context, article domain.Article) error {
	return s.artRepo.InputArticle(ctx, article)
}

func (s *syncService) InputUser(ctx context.Context, user domain.User) error {
	return s.userRepo.InputUser(ctx, user)

}

func NewSyncService(userRepo repository.UserRepository, artRepo repository.ArticleRepository) SyncService {
	return &syncService{
		userRepo: userRepo,
		artRepo:  artRepo,
	}
}
