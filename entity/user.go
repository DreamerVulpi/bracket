package entity

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
}

type UserAddRequest struct {
	Nickname string `json:"nickname"`
}

type UserEditRequest struct {
	Player User `json:"player"`
}

type UserDeleteRequest struct {
	Id int `json:"id"`
}

type UserGetRequest struct {
	Id int `json:"id"`
}

type UserAddResponse struct {
	Id int `json:"id"`
}

type UserEditResponse struct{}

type UserDeleteResponse struct{}

type UserGetResponse struct {
	Player User `json:"player"`
}
