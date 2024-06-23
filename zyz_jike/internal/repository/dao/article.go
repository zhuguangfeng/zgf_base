package dao

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"gorm.io/gorm"
)

type ArticleDao interface {
	InsertArticle(ctx context.Context, article Article) (art Article, err error)
	UpdateArticle(ctx context.Context, article Article) (err error)
	DeleteArticle(ctx context.Context, ids []int64) error
	FindArticleById(ctx context.Context, id int64) (Article, error)
}

type GormArticleDao struct {
	db *gorm.DB
}

func NewArticleDao(db *gorm.DB) ArticleDao {
	return &GormArticleDao{
		db: db,
	}
}

func (dao *GormArticleDao) InsertArticle(ctx context.Context, article Article) (art Article, err error) {
	err = dao.db.WithContext(ctx).Create(&article).Error
	return article, err
}
func (dao *GormArticleDao) UpdateArticle(ctx context.Context, article Article) (err error) {
	return dao.db.WithContext(ctx).Where("id = ?", article.ID).Updates(&article).Error
}
func (dao *GormArticleDao) DeleteArticle(ctx context.Context, ids []int64) error {
	return dao.db.WithContext(ctx).Where("id in (?)", ids).Delete(&Article{}).Error
}
func (dao *GormArticleDao) FindArticleById(ctx context.Context, id int64) (Article, error) {
	var art Article
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&art).Error
	return art, err
}

const TableNameArticle = "article"

type stringSlice []string

type Article struct {
	*Model
	Uid             int64       `gorm:"column:uid;type:bigint;index:index_uid;not null;default:0;comment:用户id" json:"uid"`                               //用户id
	Category        int8        `gorm:"column:category;type:tinyint(1);not null;comment:分类 1知友圈 2官方咨询" json:"category"`                                  //分类 1知友圈 2官方咨询
	ArticleCategory int8        `gorm:"column:article_category;type:tinyint(1);not null;default:0;comment:文章分类(知友圈和官方咨询的分类不一样)" json:"article_category"` //文章分类(知友圈和官方咨询的分类不一样) 字典news_category 1产品知识 2新闻咨询 3行业咨询
	Title           string      `gorm:"column:title;type:varchar(255);index:index_title;not null;comment:文章标题" json:"title"`                             //文章标题
	Content         string      `gorm:"column:content;type:text;not null;comment:用户ID" json:"content"`                                                   //文章内容
	RichText        string      `gorm:"column:rich_text;type:text;comment:新闻富文本" json:"rich_text"`                                                       //新闻富文本
	Pic             string      `gorm:"column:pic;type:text;not null;comment:文章首图" json:"pic"`                                                           //首图
	Pics            stringSlice `gorm:"column:pics;type:text;not null;comment:文章图片(多个)" json:"pics"`                                                     //图片
	Status          int8        `gorm:"column:status;type:tinyint(1);not null;comment:文章状态 1未发布 2发布" json:"status"`                                      //文章状态  1未发布 2发布
}

func (s *stringSlice) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, s)
}

func (s stringSlice) Value() (driver.Value, error) {
	return json.Marshal(s)
}

func (*Article) TableName() string {
	return TableNameArticle
}
