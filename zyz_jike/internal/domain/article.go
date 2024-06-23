package domain

import "time"

type Article struct {
	Id              int64     `json:"id"`
	Author          Author    `json:"author"`           //作者
	Category        int8      `json:"category"`         //分类 1知友圈 2官方咨询
	ArticleCategory int8      `json:"article_category"` //文章分类(知友圈和官方咨询的分类不一样) 字典news_category 1产品知识 2新闻咨询 3行业咨询
	Title           string    `json:"title"`            //文章标题
	Content         string    `json:"content"`          //文章内容
	RichText        string    `json:"rich_text"`        //新闻富文本
	Pic             string    `json:"pic"`              //首图
	Pics            []string  `json:"pics"`             //图片
	Status          int8      `json:"status"`           //文章状态  1未发布 2发布
	CreatedAt       time.Time ` json:"createdAt"`
}

// 文章状态
const (
	ArticleStatusUnknown     = iota //未知状态
	ArticleStatusUnPublished = 1    //为发表
	ArticlePublished         = 2    //已发表
)

type Author struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
}
