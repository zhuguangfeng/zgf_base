package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"zyz_jike/internal/domain"
	"zyz_jike/internal/errs"
	"zyz_jike/internal/repository"
)

type UserService interface {
	UserRegister(ctx context.Context, user domain.User) error
	PwdLogin(ctx *gin.Context, phone string, password string) (domain.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// 用户注册
func (s *userService) UserRegister(ctx context.Context, user domain.User) error {
	//密码加密
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return s.userRepo.CreateUser(ctx, user)
}

// 密码登录
func (s *userService) PwdLogin(ctx *gin.Context, phone string, password string) (domain.User, error) {
	user, err := s.userRepo.FindUserByPhone(ctx, phone)
	if err == errs.ErrUserNotFound {
		return domain.User{}, errs.ErrDuplicatePhone
	}
	if err != nil {
		return domain.User{}, err
	}

	// 检查密码对不对
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return domain.User{}, errs.ErrInvalidPassword
	}
	return user, nil
}
