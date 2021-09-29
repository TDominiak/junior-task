package api

import "net/http"

//TickerHandler

type TickerHandler interface {
	Start(http.ResponseWriter, *http.Request)
	Stop(http.ResponseWriter, *http.Request)

}
