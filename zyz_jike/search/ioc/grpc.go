package ioc

import (
	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
	grpc2 "google.golang.org/grpc"
	"zyz_jike/pkg/grpcx"
	"zyz_jike/pkg/logger"
	"zyz_jike/search/grpc"
)

func InitGrpcxServer(syncRpc *grpc.SyncServiceServer, searchRpc *grpc.SearchServiceServer, etcdCli *clientv3.Client, l logger.Logger) *grpcx.Server {
	type Config struct {
		Port     int    `yaml:"port"`
		EtcdAddr string `yaml:"etcdAddr"`
		EtcdTTL  int64  `yaml:"etcdTTL"`
		Name     string `yaml:"name"`
	}
	var cfg Config
	err := viper.UnmarshalKey("grpc.server", &cfg)
	if err != nil {
		panic(err)
	}
	server := grpc2.NewServer()
	syncRpc.Register(server)
	searchRpc.Register(server)
	return &grpcx.Server{
		Server:     server,
		Port:       cfg.Port,
		Name:       cfg.Name,
		L:          l,
		EtcdTTL:    cfg.EtcdTTL,
		EtcdClient: etcdCli,
	}
}
