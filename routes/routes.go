package routes

import (
	"github.com/alainmucyo/crud/controller"
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/books", controller.GetBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}", controller.ShowBook).Methods("GET")
	r.HandleFunc("/api/books", controller.StoreBook).Methods("POST")
	r.HandleFunc("/api/book/{id}", controller.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE")
	return r
}
