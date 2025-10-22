package usecase

import (
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type AuthRepo interface {
	CheckHash(nickname, password string) error
	GenerateHash(password string) (string, error)
	GenerateToken(id, secretKey string) (string, error)
}

type Auth struct {
	Repo      AuthRepo
	SecretKey string
}

type Claims struct {
	jwt.RegisteredClaims
}

func (a *Auth) CreateJWTtoken(id string) (string, error) {
	token, err := a.Repo.GenerateToken(id, a.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *Auth) CreatePasswordHash(password string) (string, error) {
	hash, err := a.Repo.GenerateHash(password)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func (a *Auth) VerifyHash(nickname, password string) error {
	err := a.Repo.CheckHash(nickname, password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
