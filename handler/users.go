package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/DreamerVulpi/bracket/jwt"
	"github.com/DreamerVulpi/bracket/usecase"
	"github.com/emersion/go-bcrypt"
	"github.com/gorilla/mux"
)

type Handler struct {
	UserUsecase usecase.User
	SetUsecase  usecase.Set
	PoolUsecase usecase.Pool
}

func readParamInt(r *http.Request, name string) (int, error) {
	vars := mux.Vars(r)
	if vars[name] == "" {
		return 0, fmt.Errorf("no %v in url string", name)
	}

	id, err := strconv.Atoi(vars[name])
	if err != nil {
		return 0, err
	}
	return id, nil
}

func readJsonRequest[T any](body io.ReadCloser) (T, error) {
	var req T
	jsonData, err := io.ReadAll(body)
	if err != nil {
		return req, fmt.Errorf("read request body: %w", err)
	}
	defer func() {
		if err := body.Close(); err != nil {
			log.Printf("close request body: %v", err)
		}
	}()

	if err := json.Unmarshal(jsonData, &req); err != nil {
		return req, fmt.Errorf("unmarshal request body: %w", err)
	}

	return req, nil
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

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := readJsonRequest[entity.UserAddRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.UserUsecase.AddUser(jsonRequest)
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

func (h *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	_, err := jwt.VerifyToken(r.Header.Get("token"))
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := readParamInt(r, "id")
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonRequest, err := readJsonRequest[entity.UserEditRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(jsonRequest.Password), 2)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonRequest.Password = string(hash)

	response, err := h.UserUsecase.EditUser(id, jsonRequest)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, err := jwt.VerifyToken(r.Header.Get("token"))
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := readParamInt(r, "id")
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	response, err := h.UserUsecase.DeleteUser(id)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	_, err := jwt.VerifyToken(r.Header.Get("token"))
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	id, err := readParamInt(r, "id")
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	response, err := h.UserUsecase.GetUser(id)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}
