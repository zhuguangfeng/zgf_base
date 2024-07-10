package domain

type Dynamic struct {
	Id        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Resources []string `json:"resources"`
}
