package jsoncontrollers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/AlyHKafoury/oaas/database"
	"github.com/AlyHKafoury/oaas/userdefinedmodels"
)

type updateRequest struct {
	DropColumn   []string `json:DropColumn`
	CreateColumn map[string]string `json:CreateColumn`
	RenameColumn map[string]string `json:RenameColumn`
  UpdateColumn map[string]string `json:UpdateColumn`
}

func ListModels(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List..!"))
}

func GetModel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET..!"))
}

func CreateModel(w http.ResponseWriter, r *http.Request) {
	var request_model userdefinedmodels.Usermodel
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	database.CreateModel()
	decoder.Decode(&request_model)
	fmt.Println("%v+", request_model)
	fmt.Println("%v+", vars)
	w.Write([]byte(request_model["aly"]))
}

func SetModels(w http.ResponseWriter, r *http.Request) {
	var request_models userdefinedmodels.Multipleusermodels
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&request_models)
	fmt.Println("%v+", request_models)
	w.Write([]byte(request_models["aly"]["bedo"]))
}

func UpdateModel(w http.ResponseWriter, r *http.Request) {
  var update_model updateRequest
  vars := mux.Vars(r)
  decoder := json.NewDecoder(r.Body)
  decoder.Decode(&update_model)

  fmt.Println("%v+", update_model)
  fmt.Println("%v+", vars)
  json.NewEncoder(w).Encode(update_model)
}

func DeleteModel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE..!"))
}