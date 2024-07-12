package repository

import (
	"context"
	"fmt"
	"webook/internal/domain"
	"webook/internal/model"
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
)

type CacheDynamicRepository struct {
	dao   dao.DynamicDao
	cache cache.DynamicCache
}

func NewDynamicRepository(dao dao.DynamicDao, cache cache.DynamicCache) DynamicRepository {
	return &CacheDynamicRepository{
		dao:   dao,
		cache: cache,
	}
}

func (repo *CacheDynamicRepository) CreateDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	res, err := repo.dao.InsertDynamic(ctx, repo.toDao(dynamic))
	if err != nil {
		return err
	}
	err = repo.cache.Set(ctx, repo.toDomain(res))
	if err != nil {
		fmt.Println(err)
		//记录日志
	}
	return nil
}

func (repo *CacheDynamicRepository) toDomain(dnm model.Dynamic) domain.Dynamic {
	return domain.Dynamic{
		Id:         dnm.BaseModel.ID,
		Title:      dnm.Title,
		Content:    dnm.Content,
		Category:   dnm.Category,
		Resources:  dnm.Resources,
		CreateTime: dnm.CreatedAt,
		UpdateTime: dnm.UpdatedAt,
	}
}

func (repo *CacheDynamicRepository) toDao(dnm domain.Dynamic) model.Dynamic {
	return model.Dynamic{
		BaseModel: model.BaseModel{
			ID: dnm.Id,
		},
		Title:     dnm.Title,
		Content:   dnm.Content,
		Resources: dnm.Resources,
		Category:  dnm.Category,
	}
}
