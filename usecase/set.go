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

func (s *Set) EditSet(id int, request entity.SetEditRequest) (entity.SetEditResponse, error) {
	_, err := s.Repo.Get(id)
	if err != nil {
		return entity.SetEditResponse{}, err
	}

	err = s.Repo.Edit(entity.Set{Id: id, Player1Id: request.Set.Player1Id, Player2Id: request.Set.Player2Id, PoolId: request.Set.PoolId})
	if err != nil {
		return entity.SetEditResponse{}, err
	}

	return entity.SetEditResponse{}, nil
}

func (s *Set) DeleteSet(id int) (entity.SetDeleteResponse, error) {
	_, err := s.Repo.Get(id)
	if err != nil {
		return entity.SetDeleteResponse{}, err
	}

	err = s.Repo.Delete(id)
	if err != nil {
		return entity.SetDeleteResponse{}, err
	}

	return entity.SetDeleteResponse{}, nil
}

func (s *Set) GetSet(id int) (entity.SetGetResponse, error) {
	set, err := s.Repo.Get(id)
	if err != nil {
		return entity.SetGetResponse{}, err
	}

	return entity.SetGetResponse{Set: set}, nil
}
