package domain

type User struct {
	Id       int64  `json:"id"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
}
