package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/DreamerVulpi/bracket/usecase"
)

type Handler struct {
	UserUsecase usecase.User
	SetUsecase  usecase.Set
}

func readRequest[T any](body io.ReadCloser) (T, error) {
	var req T
	jsonData, err := io.ReadAll(body)
	if err != nil {
		return req, fmt.Errorf("read request body: %w", err)
	}
	defer body.Close()

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
	w.Write(responseJson)
}

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	result, err := readRequest[entity.UserAddRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.UserUsecase.AddUser(result)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	player, err := readRequest[entity.UserEditRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.UserUsecase.EditUser(player)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := readRequest[entity.UserDeleteRequest](r.Body)
	if err != nil {
		log.Println(err)
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
	result, err := readRequest[entity.UserGetRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.UserUsecase.GetUser(result)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}
