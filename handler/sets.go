package handler

import (
	"log"
	"net/http"

	"github.com/DreamerVulpi/bracket/entity"
)

func (h *Handler) AddSet(w http.ResponseWriter, r *http.Request) {
	result, err := readJsonRequest[entity.SetAddRequest](r.Body)
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
	id, err := readParamInt(r, "id")
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonRequest, err := readJsonRequest[entity.SetEditRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.SetUsecase.EditSet(id, jsonRequest)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) DeleteSet(w http.ResponseWriter, r *http.Request) {
	id, err := readParamInt(r, "id")
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
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
	id, err := readParamInt(r, "id")
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	response, err := h.SetUsecase.GetSet(id)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}
