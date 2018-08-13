package main

import (
    "github.com/gorilla/mux"
    "AlyHKafoury/oaas/app/controllers"
)

func createRouter() *mux.Router {
  router := mux.NewRouter()
  router.HandleFunc("/",controllers.Index).Methods("GET")
  return router
}
