package handler

import (
	"net/http"

	"log"

	"github.com/DreamerVulpi/bracket/entity"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := readJsonRequest[entity.AuthLoginRequest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.AuthUsecase.Login(jsonRequest)
	if err != nil {
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	jsonRequest, err := readJsonRequest[entity.AuthRegisterReguest](r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	response, err := h.AuthUsecase.Register(jsonRequest)
	if err != nil {
		jsonResponse(w, entity.ErrorResponse{Error: err.Error()})
		return
	}

	jsonResponse(w, response)
}
