package main

import (
	"github.com/AlyHKafoury/oaas/jsoncontrollers"
	"github.com/gorilla/mux"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", jsoncontrollers.Index).Methods("GET")
	router.HandleFunc("/models", jsoncontrollers.ListModels).Methods("GET")
	router.HandleFunc("/models", jsoncontrollers.SetModels).Methods("POST")
	router.HandleFunc("/models/{name}", jsoncontrollers.GetModel).Methods("GET")
	router.HandleFunc("/models/{name}", jsoncontrollers.CreateModel).Methods("POST")
	router.HandleFunc("/models/{name}", jsoncontrollers.UpdateModel).Methods("PUT")
	router.HandleFunc("/models/{name}", jsoncontrollers.DeleteModel).Methods("DELETE")
	return router
}
