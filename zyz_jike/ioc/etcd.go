package ioc

import (
	"fmt"
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func InitEtcdClient() *clientv3.Client {
	type Config struct {
		Addrs []string `yaml:"addrs"`
	}
	var cfg Config
	err := viper.UnmarshalKey("etcd", &cfg)
	fmt.Println(cfg)
	if err != nil {
		panic(err)
	}
	cli, err := clientv3.NewFromURLs(cfg.Addrs)
	if err != nil {
		panic(err)
	}
	return cli
}
