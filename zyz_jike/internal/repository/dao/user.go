package dao

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
	"zyz_jike/internal/errs"
)

type UserDao interface {
	InsertUser(ctx context.Context, u User) error
	FindUserByPhone(ctx context.Context, phone string) (User, error)
	FindUserById(ctx context.Context, id int64) (User, error)
}

type GormUserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &GormUserDao{
		db: db,
	}
}

func (dao *GormUserDao) InsertUser(ctx context.Context, u User) error {
	err := dao.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {

			return errs.ErrDuplicatePhone
		}
	}
	return err
}

func (dao *GormUserDao) FindUserByPhone(ctx context.Context, phone string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("phone = ?", phone).First(&u).Error
	return u, err
}
func (dao *GormUserDao) FindUserById(ctx context.Context, id int64) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&u).Error
	return u, err
}

const TableNameUser = "user"

type User struct {
	*Model
	Phone          string       `gorm:"column:phone;type:char(11);not null;uniqueIndex:unique_phone,priority:1;comment:手机号" json:"phone"`
	Password       string       `gorm:"column:password;type:varchar(100);not null;comment:用户密码" json:"password"`
	Nickname       string       `gorm:"column:nickname;type:varchar(10);not null;default:知一招用户;comment:昵称" json:"nickname" d:"知一招用户"`
	Name           string       `gorm:"column:name;type:varchar(10);not null;default:'';comment:昵称" json:"name"`
	IdCard         string       `gorm:"column:id_card;type:varchar(18);not null;default:'';comment:身份证号码" json:"id_card"`
	Avatar         string       `gorm:"column:avatar;type:varchar(255);default:https://file.zyz.team/miniapp/my/default_avatar.jpg;comment:头像" json:"avatar" d:"https://file.zyz.team/miniapp/my/default_avatar.jpg"`
	Gender         int8         `gorm:"column:gender;type:tinyint(1);not null;default:0;comment:性别 0 未知 1 男 2 女" json:"gender"`
	Unionid        string       `gorm:"column:unionid;type:varchar(50);comment:微信统一标识" json:"unionid"`
	Openid         string       `gorm:"column:openid;type:varchar(50);comment:微信小程序渠道ID" json:"openid"`
	DateBirth      sql.NullTime `gorm:"column:date_birth;type:date;comment:出生日期" json:"date_birth"`
	TotalScore     int64        `gorm:"column:total_score;type:int;comment:总积分;default:0" json:"total_score"`
	LastLoginIP    string       `gorm:"column:last_login_ip;type:varchar(30);comment:最后登录IP" json:"last_login_ip"`
	LastLoginTime  time.Time    `gorm:"column:last_login_time;type:datetime(3);default:null;comment:最后登录时间" json:"last_login_time"`
	RegisterSource int32        `gorm:"column:register_source;type:int;comment:注册来源 1-小程序 2-App 3-web;default:0" json:"register_source"`
	Status         int8         `gorm:"column:status;type:tinyint(1);not null;default:1;comment:账户状态 1 正常 2 封号 3 注销" json:"status"`
	StopReason     string       `gorm:"column:stop_reason;type:text;comment:账户停止原因" json:"stop_reason"`
}

func (u *User) TableName() string {
	return TableNameUser
}
