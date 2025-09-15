package entity

type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}

type RequestUserAdd struct {
	Nickname string `json:"nickname"`
}

type RequestUserEdit struct {
	Player User `json:"player"`
}

type RequestUserDelete struct {
	Id string `json:"id"`
}

type RequestUserGet struct {
	Id string `json:"id"`
}

type ResponseUserAdd struct {
	Id int `json:"id"`
}

type ResponseUserEdit struct{}

type ResponseUserDelete struct{}

type ResponseUserGet struct {
	Player User `json:"player"`
}
