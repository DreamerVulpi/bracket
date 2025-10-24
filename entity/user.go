package entity

type User struct {
	Id            int    `json:"id"`
	Nickname      string `json:"nickname"`
	Password_Hash string `json:"password_hash"`
}

type UserAddRequest struct {
	Nickname      string `json:"nickname"`
	Password_Hash string `json:"password_hash"`
}

type UserEditRequest struct {
	Nickname      string `json:"nickname"`
	Password_Hash string `json:"password_hash"`
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
