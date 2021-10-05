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
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	l := log.New(os.Stdout, "tsm: ", log.LstdFlags)

	// create a new server
	s := newHttpServer(os.Getenv("APP_PORT"))

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

	// repo := repository.NewInMemortRepository() a
	repo := repository.NewMongoRepository(os.Getenv("MONGO_URL"), "device", 5)
	serivce := domain.NewDeviceService(repo)
	handler := api.NewDeviceHandler(serivce)

	router.HandleFunc("/devices", handler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/device", handler.Save).Methods(http.MethodPost)
	router.HandleFunc("/devices/{id}", handler.GetByID).Methods(http.MethodGet)
	router.Use(loggingMiddleware)

	// exchanger, err := measurementUtils.NewMeasurementStdoutExchanger(make(chan measurementUtils.Measurement))
	exchanger, err := measurementUtils.NewMeasurementRabbitExchanger()
	if err != nil {
		log.Printf("Error initialize exchanger: %s", err)
		os.Exit(1)
	}

	tickerHandler := api.NewTickerHandler(domain.NewTickerService(serivce, exchanger))
	router.HandleFunc("/start", tickerHandler.Start).Methods(http.MethodPost)
	router.HandleFunc("/stop", tickerHandler.Stop).Methods(http.MethodPost)

	err = exchanger.StartConsuming()
	if err != nil {
		log.Printf("Error start consuming: %s", err)
		os.Exit(1)
	}

	return router
}

func newHttpServer(port string) *http.Server {

	router := setupHttpRouter()
	log.Println("Starting on port", port)

	listener := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	return listener
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Url requested: %s", r.RequestURI)
		next.ServeHTTP(w, r)
		log.Println("Request finished")
	})
}
