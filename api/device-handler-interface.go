package api

import "net/http"

//DeviceHandler

type DeviceHandler interface {
	GetByID(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Save(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}
