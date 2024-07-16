package dynamic

import (
	"context"
	"github.com/IBM/sarama"
	"time"
	"webook/internal/domain"
	"webook/internal/repository"
	"webook/pkg/logger"
	"webook/pkg/saramax"
)

type DynamicConsumer struct {
	client      sarama.Client
	l           logger.Logger
	dynamicRepo repository.DynamicRepository
}

func NewDynamicConsumer(client sarama.Client, l logger.Logger, dynamicRepo repository.DynamicRepository) *DynamicConsumer {
	return &DynamicConsumer{
		client:      client,
		l:           l,
		dynamicRepo: dynamicRepo,
	}
}

func (d *DynamicConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("sync_article", d.client)
	if err != nil {
		return err
	}
	go func() {
		err := cg.Consume(
			context.Background(),
			[]string{TopicSyncDynamic},
			saramax.NewHandler(d.l, d.Consume))
		if err != nil {
			d.l.Error("退出了消费 循环异常", logger.Error(err))
		}
	}()
	return err
}

func (d *DynamicConsumer) Consume(sg *sarama.ConsumerMessage, dnc DynamicEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return d.dynamicRepo.InputDynamic(ctx, d.toDomain(dnc))
}

func (d *DynamicConsumer) toDomain(dnc DynamicEvent) domain.Dynamic {
	return domain.Dynamic{
		Id:         dnc.Id,
		Title:      dnc.Title,
		Content:    dnc.Content,
		Resources:  dnc.Resources,
		Category:   dnc.Category,
		CreateTime: dnc.CreateTime,
		UpdateTime: dnc.UpdateTime,
	}
}
