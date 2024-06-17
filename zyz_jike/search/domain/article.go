package domain

type Article struct {
	Id              int64    `json:"id"`
	Uid             int64    `json:"uid"`
	Category        int8     `json:"category"`
	ArticleCategory int8     `json:"article_category"`
	Title           string   `json:"title"`
	Content         string   `json:"content"`
	RichText        string   `json:"rich_text"`
	Pic             string   `json:"pic"`
	Pics            []string `json:"pics"`
	Status          int8     `json:"status"`
}
