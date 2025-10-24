package entity

type Auth struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type AuthLoginRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}

type AuthRegisterReguest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type AuthRegisterResponse struct {
	Id int `json:"id"`
}

type AuthTokenRequest struct {
	Id int
}

type AuthTokenResponse struct {
	State bool
}
