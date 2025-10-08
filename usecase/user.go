package usecase

import (
	"github.com/DreamerVulpi/bracket/entity"
)

type UserRepo interface {
	Add(nickname string, password string, token string) (int, error)
	Get(id int) (entity.User, error)
	GetUserByNickname(nickname string) (entity.User, error)
	Delete(id int, token string) error
	Edit(player entity.User) error
}

type User struct {
	Repo UserRepo
}

func (u *User) AddUser(request entity.UserAddRequest) (entity.UserAddResponse, error) {
	id, err := u.Repo.Add(request.Nickname, request.Password, request.JWTtoken)
	if err != nil {
		return entity.UserAddResponse{}, err
	}

	return entity.UserAddResponse{Id: id}, nil
}

func (u *User) EditUser(id int, request entity.UserEditRequest) (entity.UserEditResponse, error) {
	_, err := u.Repo.Get(id)
	if err != nil {
		return entity.UserEditResponse{}, err
	}

	err = u.Repo.Edit(entity.User{Id: id, Nickname: request.Nickname})
	if err != nil {
		return entity.UserEditResponse{}, err
	}

	return entity.UserEditResponse{}, nil
}

func (u *User) DeleteUser(id int, token string) (entity.UserDeleteResponse, error) {
	user, err := u.Repo.Get(id)
	if err != nil {
		return entity.UserDeleteResponse{}, err
	}

	// TODO: CASCADE?
	err = u.Repo.Delete(user.Id, token)
	if err != nil {
		return entity.UserDeleteResponse{}, err
	}

	return entity.UserDeleteResponse{}, nil
}

func (u *User) GetUser(id int) (entity.UserGetResponse, error) {
	user, err := u.Repo.Get(id)
	if err != nil {
		return entity.UserGetResponse{}, err
	}
	return entity.UserGetResponse{User: user}, nil
}

func (u *User) GetUserByNickname(nickname string) (entity.UserGetResponse, error) {
	user, err := u.Repo.GetUserByNickname(nickname)
	if err != nil {
		return entity.UserGetResponse{}, err
	}

	return entity.UserGetResponse{User: user}, nil
}
