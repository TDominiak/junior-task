package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TDominiak/junior-task/api"
	"github.com/TDominiak/junior-task/domain"
	"github.com/TDominiak/junior-task/domain/measurementUtils"
	"github.com/TDominiak/junior-task/repository"
	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "tsm: ", log.LstdFlags)

	// create a new server
	s := newHttpServer(getEnv("PORT", "8090"))

	// start the server
	func() {
		l.Println("Starting server")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

}

func setupHttpRouter() http.Handler {
	router := mux.NewRouter()

	repo := repository.NewInMemortRepository()
	// repo := repository.NewMongoRepository("mongodb://localhost:27017", "device", 5)
	serivce := domain.NewDeviceService(repo)
	handler := api.NewDeviceHandler(serivce)

	router.HandleFunc("/devices", handler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/device", handler.Save).Methods(http.MethodPost)
	router.HandleFunc("/devices/{id}", handler.GetByID).Methods(http.MethodGet)
	router.Use(loggingMiddleware)

	chanMeasurement := make(chan measurementUtils.Measurement)
	publisher, err := measurementUtils.GetPublisher("stdout", chanMeasurement)
	conumer := measurementUtils.MeasurementConsumer{chanMeasurement}
	go conumer.Start()

	if err != nil {
		log.Printf("Error initialize pubisher", err)
		os.Exit(1)
	}
	tickerHandler := api.NewTickerHandler(domain.NewTickerService(serivce, publisher))
	router.HandleFunc("/start", tickerHandler.Start).Methods(http.MethodPost)
	router.HandleFunc("/stop", tickerHandler.Stop).Methods(http.MethodPost)
	return router
}

func newHttpServer(port string) *http.Server {

	router := setupHttpRouter()
	log.Println("Starting on port %v", port)

	listener := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return listener
}

func getEnv(envname string, defaultValue string) string {
	env := os.Getenv(envname)

	if env != "" {
		return envname
	}

	return defaultValue
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Url requested: %s", r.RequestURI)
		next.ServeHTTP(w, r)
		log.Println("Request finished")
	})
}
