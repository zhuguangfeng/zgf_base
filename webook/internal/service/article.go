package service

import (
	"context"
	"webook/internal/domain"
)

type DynamicService interface {
	PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error
}

type dynamicService struct {
}

func NewDynamicService() DynamicService {
	return &dynamicService{}
}

func (d dynamicService) PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	//TODO implement me
	panic("implement me")
}
