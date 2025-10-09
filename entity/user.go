package entity

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	JWTtoken string `json:"JWTtoken"`
}

type UserAddRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	JWTtoken string `json:"JWTtoken"`
}

type UserEditRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	JWTtoken string `json:"JWTtoken"`
}

type UserAddResponse struct {
	Id int `json:"id"`
}

type UserEditResponse struct{}

type UserDeleteResponse struct{}

type UserGetResponse struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
}
