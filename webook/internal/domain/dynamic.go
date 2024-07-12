package domain

import "time"

type Dynamic struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Resources  []string  `json:"resources"`
	Category   int8      `json:"category"` //1图片 2视频
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
