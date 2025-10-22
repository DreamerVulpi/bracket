package handler

import (
	"net/http"

	"fmt"
	"log"

	"github.com/DreamerVulpi/bracket/entity"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := readJsonRequest[entity.AuthLoginRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if len(jsonRequest.Nickname) == 0 {
		log.Println(fmt.Errorf("don't get login"))
		jsonResponse(w, entity.ErrorResponse{Error: "don't get nickname"})
		return
	}
	if len(jsonRequest.Password) == 0 {
		log.Println(fmt.Errorf("don't get password"))
		jsonResponse(w, entity.ErrorResponse{Error: "don't get password"})
		return
	}

	err = h.AuthUsecase.VerifyHash(jsonRequest.Nickname, jsonRequest.Password)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	token, err := h.AuthUsecase.CreateJWTtoken(jsonRequest.Nickname)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, entity.AuthLoginResponse{Token: token})
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := readJsonRequest[entity.AuthRegisterReguest](r.Body)
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

	token, err := h.AuthUsecase.CreateJWTtoken(jsonRequest.Nickname)
	if err != nil {
		log.Println(err)
		return
	}

	hash, err := h.AuthUsecase.CreatePasswordHash(jsonRequest.Password)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.UserUsecase.AddUser(entity.UserAddRequest{
		Nickname: jsonRequest.Nickname,
		Password: hash,
		JWTtoken: token,
	})
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}
