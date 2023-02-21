package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	AuthRoutes(r)
	UserRoutes(r)
	PostRoutes(r)
	Imagesoutes(r)
	Hired(r)
	ProjectRoutes(r)
}
