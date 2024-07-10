package dao

import (
	"context"
	"gorm.io/gorm"
	"webook/internal/model"
)

type GormDynamicDao struct {
	db *gorm.DB
}

func NewDynamicDao(db *gorm.DB) DynamicDao {
	return &GormDynamicDao{
		db: db,
	}
}

func (dao *GormDynamicDao) InsertDynamic(ctx context.Context, dynamic model.Dynamic) (model.Dynamic, error) {
	err := dao.db.WithContext(ctx).Create(&dynamic).Error
	return dynamic, err
}
