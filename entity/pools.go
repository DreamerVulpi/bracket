package entity

type Pool struct {
	Id        int `json:"id"`
	BracketId int `json:"bracketId"`
}

type PoolAddRequest struct {
	BracketId int `json:"bracketId"`
}

type PoolEditRequest struct {
	Pool Pool `json:"pool"`
}

type PoolAddResponse struct {
	Id int `json:"id"`
}

type PoolEditResponse struct{}

type PoolDeleteResponse struct{}

type PoolGetResponse struct {
	Pool Pool `json:"pool"`
}
