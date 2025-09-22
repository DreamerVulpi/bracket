package usecase

import (
	"github.com/DreamerVulpi/bracket/entity"
)

type SetRepo interface {
	Add(player1Id, player2Id, poolId int) (int, error)
	Get(id int) (entity.Set, error)
	Delete(id int) error
	Edit(set entity.Set) error
}

type Set struct {
	Repo SetRepo
}

func (s *Set) AddSet(request entity.SetAddRequest) (entity.SetAddResponse, error) {
	id, err := s.Repo.Add(request.Player1Id, request.Player2Id, request.PoolId)
	if err != nil {
		return entity.SetAddResponse{}, err
	}
	return entity.SetAddResponse{Id: id}, nil
}

func (s *Set) EditSet(request entity.SetEditRequest) (entity.SetEditResponse, error) {
	_, err := s.Repo.Get(request.Set.Id)
	if err != nil {
		return entity.SetEditResponse{}, err
	}

	err = s.Repo.Edit(entity.Set{Id: request.Set.Id, Player1Id: request.Set.Player1Id, Player2Id: request.Set.Player2Id, PoolId: request.Set.PoolId})
	if err != nil {
		return entity.SetEditResponse{}, err
	}

	return entity.SetEditResponse{}, nil
}

func (s *Set) DeleteSet(request entity.SetDeleteRequest) (entity.SetDeleteResponse, error) {
	_, err := s.Repo.Get(request.Id)
	if err != nil {
		return entity.SetDeleteResponse{}, err
	}

	err = s.Repo.Delete(request.Id)
	if err != nil {
		return entity.SetDeleteResponse{}, err
	}

	return entity.SetDeleteResponse{}, nil
}

func (s *Set) GetSet(request entity.SetGetRequest) (entity.SetGetResponse, error) {
	set, err := s.Repo.Get(request.Id)
	if err != nil {
		return entity.SetGetResponse{}, err
	}

	return entity.SetGetResponse{Set: set}, nil
}
