package handler

import (
	"net/http"

	"fmt"
	"log"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/DreamerVulpi/bracket/jwt"
	"github.com/emersion/go-bcrypt"
)

// FIXME: Vulnerability - One user can edit another user
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := readJsonRequest[entity.UserAddRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if len(jsonRequest.Nickname) == 0 {
		log.Println(fmt.Errorf("don't get login"))
		jsonResponse(w, entity.ErrorResponse{Error: fmt.Errorf("don't get nickname").Error()})
		return
	}
	if len(jsonRequest.Password) == 0 {
		log.Println(fmt.Errorf("don't get password"))
		jsonResponse(w, entity.ErrorResponse{Error: fmt.Errorf("don't get password").Error()})
		return
	}

	response, err := h.UserUsecase.GetUserByNickname(jsonRequest.Nickname)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(response.User.Password), []byte(jsonRequest.Password))
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	w.Header().Add("token", response.User.JWTtoken)

	jsonResponse(w, response)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := readJsonRequest[entity.AuthenticationRegisterReguest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if len(jsonRequest.Nickname) == 0 {
		log.Println(fmt.Errorf("don't get login"))
		jsonResponse(w, entity.ErrorResponse{Error: fmt.Errorf("don't get nickname").Error()})
		return
	}
	if len(jsonRequest.Password) == 0 {
		log.Println(fmt.Errorf("don't get password"))
		jsonResponse(w, entity.ErrorResponse{Error: fmt.Errorf("don't get password").Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(jsonRequest.Password), 2)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonRequest.Password = string(hash)

	response, err := h.UserUsecase.AddUser(entity.UserAddRequest(jsonRequest))
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := jwt.CreateJWTtoken(jsonRequest.Nickname)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("token", token)

	jsonResponse(w, response)
}
