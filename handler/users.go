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
	use usecase.User
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

func (h *Handler) AddUser(w http.ResponseWriter, r *http.Request) {
	result, err := readRequest[entity.UserAddRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.use.AddUser(result)
	if err != nil {
		log.Println(err)
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	w.Write(responseJson)
}

func (h *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	player, err := readRequest[entity.UserEditRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.use.EditUser(player)
	if err != nil {
		log.Println(err)
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	w.Write(responseJson)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := readRequest[entity.UserDeleteRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.use.DeleteUser(id)
	if err != nil {
		log.Println(err)
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(responseJson)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	result, err := readRequest[entity.UserGetRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.use.GetUser(result)
	if err != nil {
		log.Println(err)
		return
	}

	responseJson, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(responseJson)
}
