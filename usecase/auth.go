package usecase

import (
	"fmt"
	"log"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/DreamerVulpi/bracket/pkg/jwt"
	"github.com/emersion/go-bcrypt"
)

type AuthRepo interface {
	GetHash(nickname, password string) (string, error)
}

type Auth struct {
	Repo AuthRepo
	Jwt  jwt.Jwt
	User *User
}

func (a *Auth) Register(user entity.AuthRegisterReguest) (entity.UserAddResponse, error) {
	if len(user.Nickname) == 0 {
		log.Println(fmt.Errorf("don't get nickname"))
		return entity.UserAddResponse{}, fmt.Errorf("don't get nickname")
	}
	if len(user.Password) == 0 {
		log.Println(fmt.Errorf("don't get password"))
		return entity.UserAddResponse{}, fmt.Errorf("don't get password")
	}

	hash, err := a.CreatePasswordHash(user.Password)
	if err != nil {
		log.Println(err)
		return entity.UserAddResponse{}, err
	}

	response, err := a.User.AddUser(entity.UserAddRequest{
		Nickname:      user.Nickname,
		Password_Hash: hash,
	})
	if err != nil {
		log.Println(err)
		return entity.UserAddResponse{}, err
	}

	return response, nil

}

func (a *Auth) Login(user entity.AuthLoginRequest) (entity.AuthLoginResponse, error) {
	if len(user.Nickname) == 0 {
		log.Println(fmt.Errorf("don't get nickname"))
		return entity.AuthLoginResponse{}, fmt.Errorf("don't get nickname")
	}
	if len(user.Password) == 0 {
		log.Println(fmt.Errorf("don't get password"))
		return entity.AuthLoginResponse{}, fmt.Errorf("don't get password")
	}

	err := a.VerifyHash(user.Nickname, user.Password)
	if err != nil {
		log.Println(err)
		return entity.AuthLoginResponse{}, err
	}

	token, err := a.Jwt.CreateJWTtoken(user.Nickname)
	if err != nil {
		log.Println(err)
		return entity.AuthLoginResponse{}, err
	}

	return entity.AuthLoginResponse{Token: token}, err
}

func (a *Auth) CreatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func (a *Auth) VerifyHash(nickname, password string) error {
	hash, err := a.Repo.GetHash(nickname, password)
	if err != nil {
		log.Println(err)
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
