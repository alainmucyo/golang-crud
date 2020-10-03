package routes

import (
	"github.com/alainmucyo/crud/data/books"
	"github.com/gorilla/mux"
)

type bookRouter struct {
	books.BookController
}

type BookRouter interface {
	BookRouter() *mux.Router
}

func NewBookRouter(controller books.BookController) BookRouter {
	return &bookRouter{controller}
}
func (controller *bookRouter) BookRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.GetBooks).Methods("GET")
	//r.HandleFunc("/api/book/{id}", controller.ShowBook).Methods("GET")
	r.HandleFunc("/", controller.CreateBook).Methods("POST")
	/*	r.HandleFunc("/api/book/{id}", controller.UpdateBook).Methods("PUT")
		r.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE")*/
	return r
}
