package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type model map[string]string
type models map[string]model
type updateRequest map[string]interface{}
// type updateRequest struct {
// 	dropColumn   []string `json:dropColumn`
// 	createColumn map[string]string `json:createColumn`
// 	renameColumn map[string]string `json:renameColumn`
//   updateColumn map[string]string `json:updateColumn`
// }

func ListModels(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List..!"))
}

func GetModel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GET..!"))
}

func CreateModel(w http.ResponseWriter, r *http.Request) {
	var request_model model
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&request_model)
	fmt.Println("%v+", request_model)
	fmt.Println("%v+", vars)
	w.Write([]byte(request_model["aly"]))
}

func SetModels(w http.ResponseWriter, r *http.Request) {
	var request_models models
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&request_models)
	fmt.Println("%v+", request_models)
	w.Write([]byte(request_models["aly"]["bedo"]))
}

func UpdateModel(w http.ResponseWriter, r *http.Request) {
  // test_model := updateRequest{dropColumn: []string{"g", "h", "i"},
  //                              createColumn:  map[string]string{"fix":"bex"},
  //                              renameColumn:  map[string]string{"fix":"bex"},
  //                              updateColumn:  map[string]string{"fix":"bex"}}
  var update_model updateRequest
  vars := mux.Vars(r)
  decoder := json.NewDecoder(r.Body)
  decoder.Decode(&update_model)
  // out, err := json.Marshal(test_model)
  // if err != nil {
  //     fmt.Println(err)
  //     return
  // }
  // fmt.Println("%v+", out)
  fmt.Println("%v+", update_model)
  fmt.Println("%v+", vars)
  // json.NewEncoder(w).Encode(test_model)
  //w.Write([]byte(update_model))
}

func DeleteModel(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("DELETE..!"))
}
