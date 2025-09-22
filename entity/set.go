package entity

type Set struct {
	Id        int `json:"id"`
	Player1Id int `json:"player1Id"`
	Player2Id int `json:"player2Id"`
	PoolId    int `json:"poolId"`
}

type SetAddRequest struct {
	Player1Id int `json:"player1Id"`
	Player2Id int `json:"player2Id"`
	PoolId    int `json:"poolId"`
}

type SetEditRequest struct {
	Set Set `json:"set"`
}

type SetDeleteRequest struct {
	Id int `json:"id"`
}

type SetGetRequest struct {
	Id int `json:"id"`
}

type SetAddResponse struct {
	Id int `json:"id"`
}

type SetEditResponse struct{}

type SetDeleteResponse struct{}

type SetGetResponse struct {
	Set Set `json:"set"`
}
