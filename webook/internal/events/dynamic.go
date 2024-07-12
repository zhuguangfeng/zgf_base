package events

import (
	"context"
	"github.com/IBM/sarama"
	"webook/pkg/logger"
	"webook/pkg/saramax"
)

const TopicSyncDynamic = "sync_dynamic_event"

type DynamicConsumer struct {
	client sarama.Client
	l      logger.Logger
}

func NewDynamicConsumer(client sarama.Client, l logger.Logger) *DynamicConsumer {
	return &DynamicConsumer{
		client: client,
		l:      l,
	}
}

func (d *DynamicConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("sync_article", d.client)
	if err != nil {
		return err
	}

	go func() {
		err := cg.Consume(context.Background(), []string{TopicSyncDynamic}, saramax.NewHandler(d.l, d.))
	}()
}
