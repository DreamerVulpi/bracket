package entity

type Set struct {
	Id      int  `json:"id"`
	Player1 User `json:"player1"`
	Player2 User `json:"player2"`
}
