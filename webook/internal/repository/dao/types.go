package dao

import (
	"context"
	"webook/internal/model"
)

type DynamicDao interface {
	InsertDynamic(ctx context.Context, dynamic model.Dynamic) (model.Dynamic, error)
}
