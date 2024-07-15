package service

import (
	"context"
	v1 "webook/api/dynamic/v1"
	"webook/internal/domain"
	"webook/internal/events"
	"webook/internal/repository"
	"webook/pkg/logger"
)

type DynamicService interface {
	PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error
	SearchDynamicPage(ctx context.Context, search v1.SearchDynamicListReq) ([]domain.Dynamic, error)
}

type dynamicService struct {
	repo     repository.DynamicRepository
	dncEvent events.Producer
	l        logger.Logger
}

func NewDynamicService(repo repository.DynamicRepository, dncEvent events.Producer, l logger.Logger) DynamicService {
	return &dynamicService{
		repo:     repo,
		dncEvent: dncEvent,
		l:        l,
	}
}

func (s *dynamicService) PublishDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	dnc, err := s.repo.CreateDynamic(ctx, dynamic)
	if err != nil {
		return err
	}
	dncEvent := events.DynamicEvent{
		Id:         dnc.Id,
		Title:      dnc.Title,
		Content:    dnc.Content,
		Resources:  dnc.Resources,
		Category:   dnc.Category,
		CreateTime: dnc.CreateTime,
		UpdateTime: dnc.UpdateTime,
	}
	go func() {
		err := s.dncEvent.ProducerSyncDynamicEvent(dncEvent)
		if err != nil {
			s.l.Error("发布动态 同步动态生产者 生产消息失败", logger.Field{Key: "data", Val: dncEvent}, logger.Error(err))
		}
	}()
	return nil
}

func (s *dynamicService) SearchDynamicPage(ctx context.Context, search v1.SearchDynamicListReq) ([]domain.Dynamic, error) {
	return s.repo.SearchDynamic(ctx, search.KeyWord, search.Category, search.Page, search.Size)
}
