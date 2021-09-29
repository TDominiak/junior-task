package api

import "github.com/TDominiak/junior-task/domain"

type tickerHandler struct {
	service domain.Service
}

func NewTickerHandler(service domain.Service) TickerHandler {
	return &tickerHandler{service: service}
}


func (h *deviceHandler) Save(w http.ResponseWriter, r *http.Request) {
}

func (h *deviceHandler) Stop(w http.ResponseWriter, r *http.Request) {
}
