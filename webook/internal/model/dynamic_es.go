package model

import (
	_ "embed"
	"time"
)

var (
	//go:embed dynamic_index.json
	DynamicIndex string
)

const DynamicIndexName = "dynamic"

type DynamicEs struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Resources  []string  `json:"resources"`
	Category   int8      `json:"category"` //1图片 2视频
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
