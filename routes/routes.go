package routes

import (
	bookController "github.com/alainmucyo/crud/data/books"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func RouteRegister(controller bookController.BookController) *mux.Router {

	r := mux.NewRouter()
	mount(r, "/api/", NewBookRouter(controller).BookRouter())
	return r
}
func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
