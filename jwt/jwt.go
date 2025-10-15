package jwt

import "github.com/golang-jwt/jwt/v5"

var secretKey = []byte("ps1")

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

func CreateJWTtoken(nickname string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = Claims{Username: nickname}

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
