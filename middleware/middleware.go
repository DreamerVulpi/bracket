package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"encoding/json"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/DreamerVulpi/bracket/pkg/jwt"
)

type Middleware struct {
	Jwt jwt.Jwt
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
		_, err := m.Jwt.ParseToken(tokenString)
		if err != nil {
			jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
			return
		}
		next(w, r)
	}
}
