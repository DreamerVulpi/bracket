package handler

import (
	"log"
	"net/http"

	"github.com/DreamerVulpi/bracket/entity"
)

func (h *Handler) AddSet(w http.ResponseWriter, r *http.Request) {
	result, err := readRequest[entity.SetAddRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.SetUsecase.AddSet(result)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) EditSet(w http.ResponseWriter, r *http.Request) {
	set, err := readRequest[entity.SetEditRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.SetUsecase.EditSet(set)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) DeleteSet(w http.ResponseWriter, r *http.Request) {
	id, err := readRequest[entity.SetDeleteRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.SetUsecase.DeleteSet(id)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) GetSet(w http.ResponseWriter, r *http.Request) {
	result, err := readRequest[entity.SetGetRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.SetUsecase.GetSet(result)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}
