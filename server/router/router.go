package router

import (
	"../middleware"
	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/exs", middleware.GetAllExs).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/doEx/{id}", middleware.ExCheck).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/undoEx/{id}", middleware.ExUncheck).Methods("PUT", "OPTIONS")
	return router
}
