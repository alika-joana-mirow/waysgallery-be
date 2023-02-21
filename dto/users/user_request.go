package usersdto

type UpdateUserRequest struct {
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Avatar   string `json:"avatar" form:"avatar"`
	Greeting string `json:"greeting" form:"greeting"`
	Arts     string `json:"arts" form:"arts"`
}
