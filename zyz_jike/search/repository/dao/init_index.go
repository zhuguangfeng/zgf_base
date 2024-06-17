package dao

import (
	"context"
	_ "embed"
	"github.com/olivere/elastic/v7"
	"golang.org/x/sync/errgroup"
	"time"
)

var (
	//go:embed user_index.json
	UserIndex string
	//go:embed article_index.json
	ArticleIndex string
)

func InitEs(client *elastic.Client) error {
	const timeout = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	var eg errgroup.Group
	eg.Go(func() error {
		return tryCreateIndex(ctx, client, UserIndexName, UserIndex)
	})
	eg.Go(func() error {
		return tryCreateIndex(ctx, client, ArticleIndexName, ArticleIndex)
	})
	return eg.Wait()
}

func tryCreateIndex(ctx context.Context, client *elastic.Client, idxName, idxCfg string) error {
	// 索引可能已经建好了
	ok, err := client.IndexExists(idxName).Do(ctx)
	if err != nil {
		return err
	}
	if ok {
		return nil
	}
	_, err = client.CreateIndex(idxName).Body(idxCfg).Do(ctx)
	return err
}
