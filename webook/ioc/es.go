package ioc

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"time"
)

func InitEsClient() *elastic.Client {
	type Config struct {
		Url   string `yaml:"url"`
		Sniff bool   `yaml:"sniff"`
	}
	var cfg Config
	err := viper.UnmarshalKey("es", &cfg)
	if err != nil {
		panic(fmt.Errorf("读取Es配置失败 %w", err))
	}
	const timeOut = time.Second * 10
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(cfg.Url),
		elastic.SetSniff(cfg.Sniff),
		elastic.SetHealthcheckTimeoutStartup(timeOut),
	}
	client, err := elastic.NewClient(opts...)
	if err != nil {
		panic(err)
	}
	//err = dao.InitEs(client)
	//if err != nil {
	//	panic(err)
	//}
	return client
}
