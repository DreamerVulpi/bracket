package handler

import (
	"net/http"

	"fmt"
	"log"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/emersion/go-bcrypt"
)

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

	err = bcrypt.CompareHashAndPassword([]byte(response.Password), []byte(jsonRequest.Password))
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	w.Header().Add("token", response.JWTtoken)

	jsonResponse(w, response)
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

	hash, err := bcrypt.GenerateFromPassword([]byte(jsonRequest.Password), 2)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonRequest.Password = string(hash)

	token, err := h.CreateJWTtoken(jsonRequest.Nickname)
	if err != nil {
		log.Println(err)
		return
	}

	request := entity.UserAddRequest{
		Nickname: jsonRequest.Nickname,
		Password: jsonRequest.Password,
		JWTtoken: token,
	}

	response, err := h.UserUsecase.AddUser(request)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	w.Header().Add("token", token)
	jsonResponse(w, response)
}

func (h *Handler) VerifyToken(id int, inputToken string) (bool, error) {
	log.Printf("id = %v | inputToken = %v", id, inputToken)
	response, err := h.AuthUsecase.GetUserToken(id)
	if err != nil {
		log.Println(err.Error())
		return false, err
	}

	if len(inputToken) == 0 {
		return false, fmt.Errorf("token field is empty")
	}

	if inputToken != response.Token {
		return false, fmt.Errorf("token isn't correct for this account")
	}

	return true, nil
}
