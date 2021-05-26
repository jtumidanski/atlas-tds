package rest

import (
	"atlas-tds/handlers"
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

func CreateRestService(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	go NewServer(l, ctx, wg, ProduceRoutes)
}

func ProduceRoutes(l logrus.FieldLogger) http.Handler {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/ms/tds").Subrouter()
	router.Use(CommonHeader)
	router.Handle("/docs", middleware.Redoc(middleware.RedocOpts{BasePath: "/ms/tds", SpecURL: "/ms/tds/swagger.yaml"}, nil))
	router.Handle("/swagger.yaml", http.StripPrefix("/ms/tds", http.FileServer(http.Dir("/"))))

	t := handlers.NewTopic(l)
	csRouter := router.PathPrefix("/topics").Subrouter()
	csRouter.HandleFunc("/", t.GetTopics).Methods("GET")
	csRouter.HandleFunc("/{topicId}", t.GetTopic).Methods("GET")

	return router
}
