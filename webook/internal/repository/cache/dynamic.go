package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
	"webook/internal/domain"
)

type RedisDynamicCache struct {
	redisCli   redis.Cmdable
	key        string
	expiration time.Duration
}

func NewDynamicCache(redisCli redis.Cmdable) DynamicCache {
	return &RedisDynamicCache{
		redisCli:   redisCli,
		key:        "dynamic:dynamic_",
		expiration: time.Hour * 24,
	}
}

func (cache *RedisDynamicCache) Set(ctx context.Context, dynamic domain.Dynamic) error {
	val, err := json.Marshal(dynamic)
	if err != nil {
		return err
	}
	return cache.redisCli.Set(ctx, fmt.Sprintf("%s%d", cache.key, dynamic.Id), val, cache.expiration).Err()
}
