package routes

import (
	"waysGallery/handlers"
	"waysGallery/middleware"
	"waysGallery/pkg/mysql"
	"waysGallery/repositories"

	"github.com/gorilla/mux"
)

func Hired(r *mux.Router) {
	hiredRepository := repositories.RepositoryHired(mysql.DB)
	h := handlers.HandlerHired(hiredRepository)

	r.HandleFunc("/hired", h.FindHired).Methods("GET")
	r.HandleFunc("/hired/{id}", h.GetHired).Methods("GET")
	r.HandleFunc("/hired", middleware.Auth(h.CreateHired)).Methods("POST")
	r.HandleFunc("/hired-user/{id}", h.FindHiredUser).Methods("GET")
	r.HandleFunc("/hired-user-order/{id}", h.FindHiredUserOrder).Methods("GET")
	r.HandleFunc("/hired/{id}", h.UpdateHired).Methods("Patch")
}
