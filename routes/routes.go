package routes

import (
	authorController "github.com/alainmucyo/crud/data/authors"
	bookController "github.com/alainmucyo/crud/data/books"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

func RouteRegister(controller bookController.BookController, authorController authorController.AuthorController) *mux.Router {

	r := mux.NewRouter()
	mount(r, "/book", NewBookRouter(controller).BookRouter())
	mount(r, "/author", NewAuthorRouter(authorController).AuthorRouter())
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
