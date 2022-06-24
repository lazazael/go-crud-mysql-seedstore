package routes

import (
	"github.com/gorilla/mux"
	"github.com/lazazael/go-crud-mysql-seedstore/pkg/controllers"
)

var RegisterSeedStoreRouter = func(router *mux.Router) {
	router.HandleFunc("/seed/", controllers.CreateSeed).Methods("POST")
	router.HandleFunc("/seed/", controllers.GetSeed).Methods("GET")
	router.HandleFunc("/seed/{seedID}", controllers.GetSeedById).Methods("GET")
	router.HandleFunc("/seed/{seedID}", controllers.UpdateSeed).Methods("PUT")
	router.HandleFunc("/seed/{seedID}", controllers.DeleteSeed).Methods("DELETE")
}
