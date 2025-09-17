package usecase

import (
	"github.com/DreamerVulpi/bracket/entity"
)

type UserRepo interface {
	Add(nickname string) (int, error)
	Get(id int) (entity.User, error)
	Delete(id int) error
	Edit(player entity.User) error
}

type User struct {
	repo UserRepo
}

func (u *User) AddUser(request entity.UserAddRequest) (entity.UserAddResponse, error) {
	id, err := u.repo.Add(request.Nickname)
	if err != nil {
		return entity.UserAddResponse{}, err
	}

	return entity.UserAddResponse{Id: id}, nil
}

func (u *User) EditUser(request entity.UserEditRequest) (entity.UserEditResponse, error) {
	user, err := u.repo.Get(request.Player.Id)
	if err != nil {
		return entity.UserEditResponse{}, err
	}

	err = u.repo.Edit(user)
	if err != nil {
		return entity.UserEditResponse{}, err
	}

	return entity.UserEditResponse{}, nil
}

func (u *User) DeleteUser(request entity.UserDeleteRequest) (entity.UserDeleteResponse, error) {
	user, err := u.repo.Get(request.Id)
	if err != nil {
		return entity.UserDeleteResponse{}, err
	}

	// TODO: CASCADE?
	err = u.repo.Delete(user.Id)
	if err != nil {
		return entity.UserDeleteResponse{}, err
	}

	return entity.UserDeleteResponse{}, nil
}

func (u *User) GetUser(request entity.UserGetRequest) (entity.UserGetResponse, error) {
	user, err := u.repo.Get(request.Id)
	if err != nil {
		return entity.UserGetResponse{}, err
	}

	return entity.UserGetResponse{Player: user}, nil
}
