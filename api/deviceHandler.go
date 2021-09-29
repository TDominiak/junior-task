package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TDominiak/junior-task/domain"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type deviceHandler struct {
	service domain.DeviceService
}

func NewDeviceHandler(service domain.DeviceService) DeviceHandler {
	return &deviceHandler{service: service}
}

func (h *deviceHandler) Save(w http.ResponseWriter, r *http.Request) {
	var d domain.Device
	d.ID = primitive.NewObjectID()
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err := h.service.Save(&d)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

func (h *deviceHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	device, err := h.service.GetByID(id)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if device == nil {
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(fmt.Sprintf("device with id %v not found", id)); err != nil {
			http.Error(w, fmt.Sprintf("Device with id %v not found id: %v", id, id), http.StatusInternalServerError)
			return
		}
	}

	err = json.NewEncoder(w).Encode(device)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *deviceHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	devices, err := h.service.GetAll()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(devices)
	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *deviceHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}
