package repository

import (
	"context"
	"webook/internal/domain"
)

type DynamicRepository interface {
	CreateDynamic(ctx context.Context, dynamic domain.Dynamic) error
}
