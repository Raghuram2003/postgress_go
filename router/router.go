package router

import (
	"postgress_go/middlewares"
	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router:= mux.NewRouter();
	router.HandleFunc("/api/stock/{id}",middlewares.GetStock).Methods("GET","OPTIONS")
	router.HandleFunc("/api/stock",middlewares.GetAllStocks).Methods("GET","OPTIONS")
	router.HandleFunc("/api/newStock",middlewares.CreateStock).Methods("POST","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middlewares.UpdateStock).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/stock/{id}",middlewares.DeleteStock).Methods("DELETE","OPTIONS")
	return router;
}