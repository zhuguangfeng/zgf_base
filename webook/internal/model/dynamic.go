package model

const TableNameDynamic = "dynamic"

type Dynamic struct {
	BaseModel
	Id        int64    `json:"id"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Resources []string `json:"resources"`
}
