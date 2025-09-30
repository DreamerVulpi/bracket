package entity

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type UserAddRequest struct {
	Nickname string `json:"nickname"`
}

type UserEditRequest struct {
	Nickname string `json:"nickname"`
}

type UserAddResponse struct {
	Id int `json:"id"`
}

type UserEditResponse struct{}

type UserDeleteResponse struct{}

type UserGetResponse struct {
	Player User `json:"player"`
}
