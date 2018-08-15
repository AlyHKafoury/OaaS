package main

import (
	"AlyHKafoury/oaas/app/controllers"
	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.Index).Methods("GET")
	router.HandleFunc("/models", controllers.ListModels).Methods("GET")
	router.HandleFunc("/models", controllers.SetModels).Methods("POST")
	router.HandleFunc("/models/{name}", controllers.GetModel).Methods("GET")
	router.HandleFunc("/models/{name}", controllers.CreateModel).Methods("POST")
	router.HandleFunc("/models/{name}", controllers.UpdateModel).Methods("PUT")
	router.HandleFunc("/models/{name}", controllers.DeleteModel).Methods("DELETE")
	return router
}
