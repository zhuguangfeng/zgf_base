package service

import (
	"context"
	"golang.org/x/sync/errgroup"
	"strings"
	"zyz_jike/search/domain"
	"zyz_jike/search/repository"
)

type SearchService interface {
	Search(ctx context.Context, uid int64, expression string) (domain.SearchResult, error)
}

type searchService struct {
	userRepo repository.UserRepository
	artRepo  repository.ArticleRepository
}

func (s *searchService) Search(ctx context.Context, uid int64, expression string) (domain.SearchResult, error) {
	keywords := strings.Split(expression, " ")
	var eg errgroup.Group
	var res domain.SearchResult
	eg.Go(func() error {
		users, err := s.userRepo.SearchUser(ctx, keywords)
		res.Users = users
		return err
	})
	eg.Go(func() error {
		arts, err := s.artRepo.SearchArticle(ctx, keywords)
		res.Articles = arts
		return err
	})
	return res, eg.Wait()
}

func NewSearchService(userRepo repository.UserRepository, artRepo repository.ArticleRepository) SearchService {
	return &searchService{
		userRepo: userRepo,
		artRepo:  artRepo,
	}
}
