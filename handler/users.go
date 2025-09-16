package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DreamerVulpi/bracket/entity"
	"github.com/DreamerVulpi/bracket/usecase"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Handler struct {
	Conn *pgxpool.Pool
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
	fmt.Fprintf(w, "Add")
	result, err := readRequest[entity.RequestUserAdd](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	response, err := usecase.AddUser(h.Conn, result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response.Id)
}

func (h *Handler) EditUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit")
	player, err := readRequest[entity.RequestUserEdit](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = usecase.EditUser(h.Conn, player)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete")
	id, err := readRequest[entity.RequestUserDelete](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = usecase.DeleteUser(h.Conn, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Get")
	result, err := readRequest[entity.RequestUserGet](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	player, err := usecase.GetUser(h.Conn, result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(player)
}
