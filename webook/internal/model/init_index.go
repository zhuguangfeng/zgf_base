package model

import (
	"context"
	_ "embed"
	"github.com/olivere/elastic/v7"
	"golang.org/x/sync/errgroup"
	"time"
)

func InitEs(client *elastic.Client) error {
	const timeOut = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	var eg errgroup.Group
	eg.Go(func() error {
		return tryCreateIndex(ctx, client, DynamicIndexName, DynamicIndex)
	})
	return eg.Wait()
}

func tryCreateIndex(ctx context.Context, client *elastic.Client, idxName, idxCfg string) error {
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
