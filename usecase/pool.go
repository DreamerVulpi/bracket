package usecase

import "github.com/DreamerVulpi/bracket/entity"

type PoolRepo interface {
	Add(bracket_id int) (int, error)
	Get(id int) (entity.Pool, error)
	Delete(id int) error
	Edit(pool entity.Pool) error
}

type Pool struct {
	Repo PoolRepo
}

func (p *Pool) AddPool(request entity.PoolAddRequest) (entity.PoolAddResponse, error) {
	id, err := p.Repo.Add(request.BracketId)
	if err != nil {
		return entity.PoolAddResponse{}, err
	}

	return entity.PoolAddResponse{Id: id}, nil
}

func (p *Pool) EditPool(request entity.PoolEditRequest) (entity.PoolEditResponse, error) {
	_, err := p.Repo.Get(request.Pool.Id)
	if err != nil {
		return entity.PoolEditResponse{}, err
	}

	err = p.Repo.Edit(entity.Pool{Id: request.Pool.Id, BracketId: request.Pool.BracketId})
	if err != nil {
		return entity.PoolEditResponse{}, err
	}

	return entity.PoolEditResponse{}, nil
}

func (p *Pool) DeletePool(request entity.PoolDeleteRequest) (entity.PoolDeleteResponse, error) {
	pool, err := p.Repo.Get(request.Id)
	if err != nil {
		return entity.PoolDeleteResponse{}, err
	}

	// TODO: CASCADE?
	err = p.Repo.Delete(pool.Id)
	if err != nil {
		return entity.PoolDeleteResponse{}, err
	}

	return entity.PoolDeleteResponse{}, nil
}

func (p *Pool) GetPool(request entity.PoolGetRequest) (entity.PoolGetResponse, error) {
	pool, err := p.Repo.Get(request.Id)
	if err != nil {
		return entity.PoolGetResponse{}, err
	}

	return entity.PoolGetResponse{Pool: pool}, nil
}
