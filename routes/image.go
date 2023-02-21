package routes

import (
	"waysGallery/handlers"
	"waysGallery/middleware"
	"waysGallery/pkg/mysql"
	"waysGallery/repositories"

	"github.com/gorilla/mux"
)

func Imagesoutes(r *mux.Router) {
	imagesRepository := repositories.RepositoryImages(mysql.DB)
	h := handlers.HandlerImages(imagesRepository)

	r.HandleFunc("/images", h.FindImages).Methods("GET")
	r.HandleFunc("/image/{id}", h.GetImages).Methods("GET")
	r.HandleFunc("/image", middleware.Auth(h.CreateImages)).Methods("POST")
	// r.HandleFunc("/post/{id}", middleware.Auth(h.UpdatePost)).Methods("PATCH")
	// r.HandleFunc("/post/{id}", middleware.Auth(h.Deleteimages)).Methods("DELETE")
}
