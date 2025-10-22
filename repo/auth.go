package repo

import (
	"fmt"
	"log"

	"context"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/emersion/go-bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Auth struct {
	Conn *pgxpool.Pool
}

type Claims struct {
	jwt.RegisteredClaims
}

func (a *Auth) CheckHash(nickname, password string) error {
	const sql = "SELECT u.password FROM users u WHERE nickname = $1"

	var user entity.User

	err := a.Conn.QueryRow(context.Background(), sql, nickname).Scan(&user.Password)
	if err != nil {
		return fmt.Errorf("unable to get from database using nickname, %w", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (a *Auth) GenerateHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 2)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(hash), nil
}

func (a *Auth) GenerateToken(id, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = Claims{jwt.RegisteredClaims{ID: id}}
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
