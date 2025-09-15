package entity

type Tournament struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Events []Event `json:"events"`
}
