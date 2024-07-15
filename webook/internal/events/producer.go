package events

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"time"
)

const TopicSyncDynamic = "sync_dynamic_event"

type SaramaProducer struct {
	producer sarama.SyncProducer
}

type DynamicEvent struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Resources  []string  `json:"resources"`
	Category   int8      `json:"category"` //1图片 2视频
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func NewSaramaProducer(producer sarama.SyncProducer) Producer {
	return &SaramaProducer{
		producer: producer,
	}
}

// 发送同步文章消息
func (s *SaramaProducer) ProducerSyncDynamicEvent(dnc DynamicEvent) error {
	val, err := json.Marshal(&dnc)
	if err != nil {
		return err
	}
	_, _, err = s.producer.SendMessage(&sarama.ProducerMessage{
		Topic: TopicSyncDynamic,
		Value: sarama.StringEncoder(val),
	})
	return err
}
