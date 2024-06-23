package article

import (
	"encoding/json"
	"github.com/IBM/sarama"
)

const TopicSyncArticle = "sync_article_event"

type Producer interface {
	ProduceSyncArticleEvent(evt ArticleEvent) error
}
type SaramaSyncArticleProducer struct {
	producer sarama.SyncProducer
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

func NewSaramaSyncArticleProducer(producer sarama.SyncProducer) Producer {
	return &SaramaSyncArticleProducer{producer: producer}
}
func (s *SaramaSyncArticleProducer) ProduceSyncArticleEvent(evt ArticleEvent) error {
	val, err := json.Marshal(&evt)
	if err != nil {
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: TopicSyncArticle,
		Value: sarama.StringEncoder(val),
	})
	return err
}
