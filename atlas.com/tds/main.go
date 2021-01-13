package main

import (
	"atlas-wrg/handlers"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	handleRequests()
}

func handleRequests() {
	l := log.New(os.Stdout, "tds ", log.LstdFlags)

	router := mux.NewRouter().StrictSlash(true).PathPrefix("/ms/tds").Subrouter()
	router.Use(commonHeader)
	router.Handle("/docs", middleware.Redoc(middleware.RedocOpts{BasePath: "/ms/tds", SpecURL: "/ms/tds/swagger.yaml"}, nil))
	router.Handle("/swagger.yaml", http.StripPrefix("/ms/tds", http.FileServer(http.Dir("/"))))

	t := handlers.NewTopic(l)
	csRouter := router.PathPrefix("/topics").Subrouter()
	csRouter.HandleFunc("/", t.GetTopics).Methods("GET")
	csRouter.HandleFunc("/{topicId}", t.GetTopic).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func commonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
