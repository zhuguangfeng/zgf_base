package v1

import (
	"time"
	"webook/api"
)

type Dynamic struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	Resources  []string  `json:"resources"`
	Category   int8      `json:"category"` //1图片 2视频
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

// 发布动态
type PublishDynamicReq struct {
	Title     string   `json:"title"`     //标题
	Content   string   `json:"content"`   //内容
	Category  int8     `json:"category"`  //分类
	Resources []string `json:"resources"` //资源(图片or视频)
}
type PublishDynamicRes struct {
}

// 动态搜索
type SearchDynamicListReq struct {
	api.PageReq
	KeyWord  string `json:"keyword"`
	Category int8   `json:"category"`
}

type SearchDynamicListRes struct {
	List       []Dynamic   `json:"list"`
	Pagination api.PageRes `json:"pagination"`
}
