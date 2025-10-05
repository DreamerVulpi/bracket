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
}

type UserEditRequest struct {
	Login    string `json:"login"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	JWTtoken string `json:"JWTtoken"`
}

type UserGetRequest struct {
	JWTtoken string `json:"JWTtoken"`
}

type UserDeleteRequest struct {
	JWTtoken string `json:"JWTtoken"`
}

type UserAddResponse struct {
	Id int `json:"id"`
}

type UserEditResponse struct{}

type UserDeleteResponse struct{}

type UserGetResponse struct {
	User User `json:"user"`
}
