package routes

import (
	"waysGallery/handlers"
	"waysGallery/middleware"
	"waysGallery/pkg/mysql"
	"waysGallery/repositories"

	"github.com/gorilla/mux"
)

func PostRoutes(r *mux.Router) {
	postRepository := repositories.RepositoryPost(mysql.DB)
	h := handlers.HandlerPost(postRepository)

	r.HandleFunc("/post", h.FindPost).Methods("GET")
	r.HandleFunc("/post/{id}", h.GetPost).Methods("GET")
	r.HandleFunc("/post", middleware.Auth(h.CreatePost)).Methods("POST")
	r.HandleFunc("/post/{id}", middleware.Auth(h.UpdatePost)).Methods("PATCH")
	r.HandleFunc("/post/{id}", middleware.Auth(h.DeletePost)).Methods("DELETE")
}
