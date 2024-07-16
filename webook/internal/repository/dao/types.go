package dao

import (
	"context"
	"webook/internal/model"
)

type DynamicDao interface {
	InsertDynamic(ctx context.Context, dynamic model.Dynamic) (model.Dynamic, error)
}

type DynamicEsDao interface {
	InputDynamic(ctx context.Context, dynamic model.DynamicEs) error
	SearchDynamic(ctx context.Context, keyword string, category int8, page, size int) ([]model.DynamicEs, error)
}
