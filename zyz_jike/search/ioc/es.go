package ioc

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"time"
	"zyz_jike/search/repository/dao"
)

func InitEsClient() *elastic.Client {
	type Config struct {
		Url   string `json:"url"`
		Sniff bool   `json:"sniff"`
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
	err = dao.InitEs(client)
	if err != nil {
		panic(err)
	}
	return client
}
