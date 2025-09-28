package handler

import (
	"log"
	"net/http"

	"github.com/DreamerVulpi/bracket/entity"
)

func (h *Handler) AddPool(w http.ResponseWriter, r *http.Request) {
	result, err := readJsonRequest[entity.PoolAddRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.PoolUsecase.AddPool(result)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) EditPool(w http.ResponseWriter, r *http.Request) {
	id, err := readParamIdRequest(r)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonRequest, err := readJsonRequest[entity.PoolEditRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.PoolUsecase.EditPool(id, jsonRequest)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) DeletePool(w http.ResponseWriter, r *http.Request) {
	id, err := readParamIdRequest(r)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	response, err := h.PoolUsecase.DeletePool(id)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) GetPool(w http.ResponseWriter, r *http.Request) {
	id, err := readParamIdRequest(r)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	response, err := h.PoolUsecase.GetPool(id)
	if err != nil {
		log.Println(err)
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}
