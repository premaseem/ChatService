//package main
//
//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"net/http"
//	"encoding/json"
//)
//
//func Index(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome")
//}
//
//func TodoIndex(w http.ResponseWriter, r *http.Request) {
//	todos := Todos{
//		Todo{Name: "Write presentation"},
//		Todo{Name: "Host meetup"},
//	}
//
//	json.NewEncoder(w).Encode(todos)
//}
//func TodoShow(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	todoId := vars["todoId"]
//	fmt.Fprintf(w, "Todo Index!" ,todoId)
//}
