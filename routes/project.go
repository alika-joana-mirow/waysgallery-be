package routes

import (
	"waysGallery/handlers"
	"waysGallery/middleware"
	"waysGallery/pkg/mysql"
	"waysGallery/repositories"

	"github.com/gorilla/mux"
)

func ProjectRoutes(r *mux.Router) {
	projectRepository := repositories.RepositoryProject(mysql.DB)
	h := handlers.HandlerProject(projectRepository)

	r.HandleFunc("/project", h.FindProject).Methods("GET")
	r.HandleFunc("/project/{id}", h.GetProject).Methods("GET")
	r.HandleFunc("/project-hired/{id}", h.FindHiredById).Methods("GET")
	r.HandleFunc("/project", middleware.Auth(h.CreateProject)).Methods("POST")
}
