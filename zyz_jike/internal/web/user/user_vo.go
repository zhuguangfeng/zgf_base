package user

type UserRegisterReq struct {
	Phone      string `json:"phone"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
