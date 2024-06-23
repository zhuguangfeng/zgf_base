package events

import (
	"context"
	"github.com/IBM/sarama"
	"time"
	"zyz_jike/pkg/logger"
	"zyz_jike/pkg/saramax"
	"zyz_jike/search/domain"
	"zyz_jike/search/service"
)

const TopicSyncArticle = "sync_article_event"

type ArticleConsumer struct {
	syncSvc service.SyncService
	client  sarama.Client
	l       logger.Logger
}

type ArticleEvent struct {
	Id              int64    `json:"id"`
	Uid             int64    `json:"uid"`
	Category        int8     `json:"category"`
	ArticleCategory int8     `json:"article_category"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	RichText        string   `json:"rich_text"`
	Pic             string   `json:"pic"`
	Pics            []string `json:"pics"`
	Status          int8     `json:"status"`
	CreatedAt       string   `json:"created_at"`
}

func NewArticleConsumer(s service.SyncService, client sarama.Client, l logger.Logger) *ArticleConsumer {
	return &ArticleConsumer{
		syncSvc: s,
		client:  client,
		l:       l,
	}
}

func (a *ArticleConsumer) Start() error {
	cg, err := sarama.NewConsumerGroupFromClient("sync_article", a.client)
	if err != nil {
		return err
	}

	go func() {
		err := cg.Consume(context.Background(), []string{TopicSyncArticle}, saramax.NewHandler[ArticleEvent](a.l, a.Consume))
		if err != nil {
			a.l.Error("退出了消费循环异常", logger.Error(err))
		}
	}()
	return err
}

func (a *ArticleConsumer) Consume(sg *sarama.ConsumerMessage, art ArticleEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	return a.syncSvc.InputArticle(ctx, a.toDomain(art))
}

func (a *ArticleConsumer) toDomain(art ArticleEvent) domain.Article {
	return domain.Article{
		Id:              art.Id,
		Uid:             art.Uid,
		Category:        int8(art.Category),
		ArticleCategory: int8(art.ArticleCategory),
		Title:           art.Title,
		Content:         art.Content,
		RichText:        art.RichText,
		Pic:             art.Pic,
		Pics:            art.Pics,
		Status:          int8(art.Status),
	}
}
