package entity

type Pool struct {
	Id   string `json:"id"`
	Sets []Set  `json:"sets"`
}
