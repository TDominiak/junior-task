package main

import (
	"log"
	"net/http"
	"os"

	"github.com/TDominiak/junior-task/api"
	"github.com/TDominiak/junior-task/domain"
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

	return router
}

func newHttpServer(port string) *http.Server {

	router := setupHttpRouter()

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

// func idMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		id := primitive.NewObjectID()
// 		r = r.WithContext(context.WithValue(r.Context(), "id", id))
// 		next.ServeHTTP(w, r)
// 	})
// }

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Url requested: %s", r.RequestURI)
		next.ServeHTTP(w, r)
		log.Println("Request finished")
	})
}

