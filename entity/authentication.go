package entity

type Authentication struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type AuthenticationLoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type AuthenticationLoginResponse struct{}

type AuthenticationRegisterReguest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type AuthenticationRegisterResponse struct {
	Id int `json:"id"`
}
