package usecase

import "github.com/DreamerVulpi/bracket/entity"

type AuthRepo interface {
	GetToken(id int) (string, error)
}

type Auth struct {
	Repo AuthRepo
}

func (a *Auth) GetUserToken(id int) (entity.AuthTokenResponse, error) {
	token, err := a.Repo.GetToken(id)
	if err != nil {
		return entity.AuthTokenResponse{}, err
	}
	return entity.AuthTokenResponse{Token: token}, nil
}
