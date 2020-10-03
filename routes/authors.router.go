package routes

import (
	"github.com/alainmucyo/crud/data/authors"
	"github.com/gorilla/mux"
)

type authorRouter struct {
	authors.AuthorController
}

func (controller *authorRouter) AuthorRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/", controller.Store).Methods("POST")
	return r
}

type AuthorRouter interface {
	AuthorRouter() *mux.Router
}

func NewAuthorRouter(controller authors.AuthorController) AuthorRouter {
	return &authorRouter{controller}
}
