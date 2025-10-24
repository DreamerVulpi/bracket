package entity

type User struct {
	Id           int    `json:"id"`
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"-"`
}

type UserAddRequest struct {
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"-"`
}

type UserEditRequest struct {
	Nickname     string `json:"nickname"`
	PasswordHash string `json:"-"`
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
