package dao

import (
	"context"
	"fmt"
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
	fmt.Println(dynamic.CreatedAt)
	fmt.Println(dynamic.UpdatedAt)
	return dynamic, err
}
