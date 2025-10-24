package jwt

import (
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	SecretKey string
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Jwt) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		log.Printf("Token parse error: %v", err)
		return nil, fmt.Errorf("incorrect token: parsing failed")
	}

	if !token.Valid {
		log.Printf("Token incorrect error: %v", err)
		return nil, fmt.Errorf("incorrect token: not valid")
	}

	return token, nil
}

func (j *Jwt) CreateJWTtoken(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = Claims{jwt.RegisteredClaims{ID: id}}
	tokenString, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
