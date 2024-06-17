package repository

import (
	"context"
	"github.com/ecodeclub/ekit/slice"
	"zyz_jike/search/domain"
	"zyz_jike/search/repository/dao"
)

type userRepository struct {
	dao dao.UserDao
}

func (repo *userRepository) InputUser(ctx context.Context, msg domain.User) error {
	return repo.dao.InputUser(ctx, dao.User{
		Id:       msg.Id,
		Nickname: msg.Nickname,
		Phone:    msg.Phone,
	})
}

func (repo *userRepository) SearchUser(ctx context.Context, keywords []string) ([]domain.User, error) {
	users, err := repo.dao.Search(ctx, keywords)
	if err != nil {
		return nil, err
	}
	return slice.Map(users, func(idx int, src dao.User) domain.User {
		return domain.User{
			Id:       src.Id,
			Nickname: src.Nickname,
			Phone:    src.Phone,
		}
	}), nil
}

func NewUserRepository(d dao.UserDao) UserRepository {
	return &userRepository{
		dao: d,
	}
}
