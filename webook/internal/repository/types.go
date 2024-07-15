package repository

import (
	"context"
	"webook/internal/domain"
)

type DynamicRepository interface {
	CreateDynamic(ctx context.Context, dynamic domain.Dynamic) (domain.Dynamic, error)
	InputDynamic(ctx context.Context, dynamic domain.Dynamic) error
	SearchDynamic(ctx context.Context, keyword string, category int8, page, size int) ([]domain.Dynamic, error)
}
