package cache

import (
	"context"
	"webook/internal/domain"
)

type DynamicCache interface {
	Set(ctx context.Context, dynamic domain.Dynamic) error
}
