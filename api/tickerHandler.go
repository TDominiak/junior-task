package api

import (
	"log"
	"net/http"

	"github.com/TDominiak/junior-task/domain"
	"gopkg.in/square/go-jose.v2/json"
)

type tickerHandler struct {
	service domain.TickerService
}

func NewTickerHandler(service domain.TickerService) TickerHandler {
	return &tickerHandler{service: service}
}

func (t *tickerHandler) Start(w http.ResponseWriter, r *http.Request) {
	if err := t.service.Start(); err != nil {
		log.Print(err)
		http.Error(w, "failed to start ticker", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Ticker start"); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func (t *tickerHandler) Stop(w http.ResponseWriter, r *http.Request) {
	if err := t.service.Stop(); err != nil {
		log.Print(err)
		http.Error(w, "failed to stop ticker", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode("Ticker stop"); err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
