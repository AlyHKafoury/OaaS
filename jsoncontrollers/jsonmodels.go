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
	var requestModel userdefinedmodels.UserModel
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	database.CreateModel()
	decoder.Decode(&requestModel)
	fmt.Println("%v+", requestModel)
	fmt.Println("%v+", vars)
	w.Write([]byte(requestModel["aly"]))
}

func SetModels(w http.ResponseWriter, r *http.Request) {
	var requestModels userdefinedmodels.MultipleUserModels
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&requestModels)
	fmt.Println("%v+", requestModels)
	w.Write([]byte(requestModels["aly"]["bedo"]))
}

func UpdateModel(w http.ResponseWriter, r *http.Request) {
  var updateModel updateRequest
  vars := mux.Vars(r)
  decoder := json.NewDecoder(r.Body)
  decoder.Decode(&updateModel)

  fmt.Println("%v+", updateModel)
  fmt.Println("%v+", vars)
  json.NewEncoder(w).Encode(updateModel)
}

func DeleteModel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE..!"))
}
