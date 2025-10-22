package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/golang-jwt/jwt/v5"
)

type Middleware struct {
	SecretKey string
}

func (m *Middleware) parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(m.SecretKey), nil
	})
}

func jsonResponse(w http.ResponseWriter, response any) {
	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(responseJson)
	if err != nil {
		log.Println(err)
		return
	}
}

func (m *Middleware) Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			jsonResponse(w, entity.ErrorResponse{Error: fmt.Errorf("no token").Error()})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := m.parseToken(tokenString)
		if err != nil {
			log.Printf("Token parse error: %v", err)
			jsonResponse(w, entity.ErrorResponse{Error: fmt.Errorf("incorrect token: parsing failed").Error()})
			return
		}

		if !token.Valid {
			jsonResponse(w, entity.ErrorResponse{Error: fmt.Errorf("incorrect token: not valid").Error()})
			return
		}

		next(w, r)
	}
}
