package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/DreamerVulpi/bracket/entity"
)

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

func AddHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Add")
	result, err := readRequest[entity.RequestUserAdd](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Edit")
	result, err := readRequest[entity.RequestUserEdit](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delete")
	result, err := readRequest[entity.RequestUserDelete](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "List")
	result, err := readRequest[entity.RequestUserGet](r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
