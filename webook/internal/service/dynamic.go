package service

import (
	"context"
	"webook/internal/domain"
	"webook/internal/repository"
)

type DynamicService interface {
	PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error
}

type dynamicService struct {
	repo repository.DynamicRepository
}

func NewDynamicService(repo repository.DynamicRepository) DynamicService {
	return &dynamicService{
		repo: repo,
	}
}

func (s *dynamicService) PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	return s.repo.CreateDynamic(ctx, dynamic)
}
