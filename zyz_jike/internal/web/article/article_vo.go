package article

type PublishArticleReq struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Status  int8     `json:"status"`
	Pics    []string `json:"pics"`
}

type ArticleVo struct {
	Id       int64    `json:"id"`
	Title    string   `json:"title"`
	Content  string   `json:"content"`
	Uid      int64    `json:"uid"`
	Nickname string   `json:"nickname"`
	Pic      string   `json:"pic"`
	Pics     []string `json:"pics"`
}
