package ioc

import (
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
	"zyz_jike/search/events"
)

func InitKafka() sarama.Client {
	type Config struct {
		Addr []string `yaml:"addr"`
	}
	saramaCfg := sarama.NewConfig()
	saramaCfg.Producer.Return.Successes = true
	var cfg Config
	err := viper.UnmarshalKey("kafka", &cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
	client, err := sarama.NewClient(cfg.Addr, saramaCfg)
	if err != nil {
		panic(err)
	}
	return client
}

func NewConsumers(articleConsumer *events.ArticleConsumer) []events.Consumer {
	return []events.Consumer{
		articleConsumer,
	}
}
