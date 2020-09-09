package routes

import (
	"github.com/alainmucyo/crud/data/books"
	"github.com/gorilla/mux"
)

func BookRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", books.GetBooks).Methods("GET")
	//r.HandleFunc("/api/book/{id}", controller.ShowBook).Methods("GET")
	r.HandleFunc("/books", books.CreateBook).Methods("POST")
	/*	r.HandleFunc("/api/book/{id}", controller.UpdateBook).Methods("PUT")
		r.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE")*/
	return r
}
/*var bookRoutes = []Route{
	Route{
		URI:     "/books",
		Method:  "GET",
		Handler: controller.GetBooks,
	},
	Route{
		URI: "/api/books",
		Method: "POST",
		Handler: controller.CreateBook,
	},
}
*/