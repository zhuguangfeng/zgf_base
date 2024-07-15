package repository

import (
	"context"
	"webook/internal/domain"
	"webook/internal/model"
	"webook/internal/repository/cache"
	"webook/internal/repository/dao"
	"webook/pkg/logger"
)

type CacheDynamicRepository struct {
	dncDao   dao.DynamicDao
	dncEsDao dao.DynamicEsDao
	cache    cache.DynamicCache
	l        logger.Logger
}

func NewDynamicRepository(dncDao dao.DynamicDao, dncEsDao dao.DynamicEsDao, cache cache.DynamicCache, l logger.Logger) DynamicRepository {
	return &CacheDynamicRepository{
		dncDao:   dncDao,
		dncEsDao: dncEsDao,
		cache:    cache,
		l:        l,
	}
}

func (repo *CacheDynamicRepository) CreateDynamic(ctx context.Context, dynamic domain.Dynamic) (domain.Dynamic, error) {
	res, err := repo.dncDao.InsertDynamic(ctx, repo.toDao(dynamic))
	if err != nil {
		return domain.Dynamic{}, err
	}
	err = repo.cache.Set(ctx, repo.toDomain(res))
	if err != nil {
		repo.l.Error("发布动态 写入缓存失败", logger.Field{Key: "data", Val: res}, logger.Error(err))
		//记录日志
	}
	return repo.toDomain(res), nil
}

func (repo *CacheDynamicRepository) InputDynamic(ctx context.Context, dynamic domain.Dynamic) error {
	return repo.dncEsDao.InputDynamic(ctx, repo.toDao(dynamic))
}

func (repo *CacheDynamicRepository) SearchDynamic(ctx context.Context, keyword string, category int8, page, size int) ([]domain.Dynamic, error) {
	dncs, err := repo.dncEsDao.SearchDynamic(ctx, keyword, category, page, size)
	if err != nil {
		return nil, err
	}
	var res = make([]domain.Dynamic, 0, len(dncs))
	if len(dncs) > 0 {
		for _, dnc := range dncs {
			res = append(res, repo.esToDomain(dnc))
		}
	}
	return res, nil
}

func (repo *CacheDynamicRepository) esToDomain(dnm model.DynamicEs) domain.Dynamic {
	return domain.Dynamic{
		Id:         dnm.Id,
		Title:      dnm.Title,
		Content:    dnm.Content,
		Category:   dnm.Category,
		Resources:  dnm.Resources,
		CreateTime: dnm.CreateTime,
		UpdateTime: dnm.CreateTime,
	}
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
