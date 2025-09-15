package entity

type Bracket struct {
	Id    string `json:"id"`
	Pools []Pool `json:"pools"`
}
