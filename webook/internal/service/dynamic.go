package service

import (
	"context"
	"webook/internal/domain"
	"webook/internal/repository"
	"webook/pkg/logger"
)

type DynamicService interface {
	PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error
}

type dynamicService struct {
	repo repository.DynamicRepository
	l    logger.Logger
}

func NewDynamicService(repo repository.DynamicRepository, l logger.Logger) DynamicService {
	return &dynamicService{
		repo: repo,
		l:    l,
	}
}

func (s *dynamicService) PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	return s.repo.CreateDynamic(ctx, dynamic)

}
