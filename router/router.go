package router

import (
	"encoding/json"
	"github.com/alainmucyo/crud/entity"
	"github.com/gorilla/mux"
	"net/http"
)

func Register() *mux.Router {
	
	r := mux.NewRouter()
	r.HandleFunc("/api/books", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		var book entity.Book
		_ = json.NewEncoder(writer).Encode(book)

	}).Methods("GET")
	return r
}
