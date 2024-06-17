package domain

import "time"

type User struct {
	Id             int64     `json:"id"`
	Phone          string    `json:"phone"`
	Password       string    `json:"password"`
	Nickname       string    ` json:"nickname" d:"知一招用户"`
	Name           string    ` json:"name"`
	IdCard         string    `json:"id_card"`
	Avatar         string    `json:"avatar" d:"https://file.zyz.team/miniapp/my/default_avatar.jpg"`
	Gender         int8      `json:"gender"` //性别 0 未知 1 男 2 女
	Unionid        string    `json:"unionid"`
	Openid         string    `json:"openid"`
	DateBirth      time.Time `json:"date_birth"`
	TotalScore     int64     `json:"total_score"`
	LastLoginIP    string    `json:"last_login_ip"`
	LastLoginTime  time.Time `json:"last_login_time"`
	RegisterSource int32     `json:"register_source"` //注册来源 1-小程序 2-App 3-web
	Status         int8      `json:"status"`          //账户状态 1 正常 2 封号 3 注销
	StopReason     string    `json:"stop_reason"`
}
