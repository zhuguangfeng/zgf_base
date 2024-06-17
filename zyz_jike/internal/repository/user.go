package repository

import (
	"context"
	"database/sql"
	"time"
	"zyz_jike/internal/domain"
	"zyz_jike/internal/repository/dao"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.User) error
	FindUserByPhone(ctx context.Context, phone string) (domain.User, error)
	FindUserById(ctx context.Context, id int64) (domain.User, error)
}
type CacheUserRepository struct {
	dao dao.UserDao
}

func NewUserRepository(dao dao.UserDao) UserRepository {
	return &CacheUserRepository{
		dao: dao,
	}
}

// 创建用户
func (c *CacheUserRepository) CreateUser(ctx context.Context, user domain.User) error {
	return c.dao.InsertUser(ctx, c.toEntity(user))
}

func (c *CacheUserRepository) FindUserByPhone(ctx context.Context, phone string) (domain.User, error) {
	daoU, err := c.dao.FindUserByPhone(ctx, phone)
	return c.toDomain(daoU), err
}

func (c *CacheUserRepository) FindUserById(ctx context.Context, id int64) (domain.User, error) {
	daoU, err := c.dao.FindUserById(ctx, id)
	return c.toDomain(daoU), err
}

func (c *CacheUserRepository) toEntity(user domain.User) dao.User {
	return dao.User{
		Phone:          user.Phone,
		Password:       user.Password,
		Nickname:       user.Nickname,
		Name:           user.Name,
		IdCard:         user.IdCard,
		Avatar:         user.Avatar,
		Gender:         user.Gender,
		Unionid:        user.Unionid,
		Openid:         user.Openid,
		DateBirth:      sql.NullTime{Time: user.DateBirth, Valid: user.DateBirth != time.Time{}},
		TotalScore:     user.TotalScore,
		LastLoginIP:    user.LastLoginIP,
		LastLoginTime:  user.LastLoginTime,
		RegisterSource: user.RegisterSource,
		Status:         user.Status,
		StopReason:     user.StopReason,
	}
}

func (c *CacheUserRepository) toDomain(u dao.User) domain.User {
	return domain.User{
		Id:             u.ID,
		Phone:          u.Phone,
		Password:       u.Password,
		Nickname:       u.Nickname,
		Name:           u.Name,
		IdCard:         u.IdCard,
		Avatar:         u.Avatar,
		Gender:         u.Gender, //性别 0 未知 1 男 2 女
		Unionid:        u.Unionid,
		Openid:         u.Openid,
		DateBirth:      u.DateBirth.Time,
		TotalScore:     u.TotalScore,
		LastLoginIP:    u.LastLoginIP,
		LastLoginTime:  u.LastLoginTime,
		RegisterSource: u.RegisterSource, //注册来源 1-小程序 2-App 3-web
		Status:         u.Status,         //账户状态 1 正常 2 封号 3 注销
		StopReason:     u.StopReason,
	}

}
