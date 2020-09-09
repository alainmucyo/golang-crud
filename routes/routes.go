package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)
type Route struct {
	URI          string
	Method       string
	Handler      func(writer http.ResponseWriter, request *http.Request)
}
/*func load() []Route{
	routes := bookRoutes
	return routes
}*/
func Register() *mux.Router {
	r := mux.NewRouter()
	/*for _,route:= range load(){
		r.HandleFunc(route.URI,route.Handler).Methods(route.Method)
	}*/
	mount(r,"/api/",BookRouter())

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

