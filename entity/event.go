package entity

type Event struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Brackets []Bracket `json:"brackets"`
}
