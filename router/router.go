package router

import (
	"github.com/gorilla/mux"

	"github.com/vviia/elepantskids/middleware"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	// r seperti router
	r.HandleFunc("/api/stock/{id}", middleware.GetStock).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/stock", middleware.GetAllStock).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/newstock", middleware.CreateStock).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/updatestock/{id}", middleware.UpdateStock).Methods("PUT", "OPTIONS")
	r.HandleFunc("/api/deletestock/{id}", middleware.DeleteStock).Methods("DELETE", "OPTIONS")
	return r
}
