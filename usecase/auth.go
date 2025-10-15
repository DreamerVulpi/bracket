package usecase

import "github.com/DreamerVulpi/bracket/entity"

type AuthRepo interface {
	CheckToken(token string) (bool, error)
}

type Auth struct {
	Repo AuthRepo
}

func (a *Auth) CheckTokenFromDb(token string) (entity.AuthTokenResponse, error) {
	state, err := a.Repo.CheckToken(token)
	if err != nil {
		return entity.AuthTokenResponse{}, err
	}
	return entity.AuthTokenResponse{State: state}, nil
}
