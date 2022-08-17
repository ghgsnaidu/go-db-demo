package routes

import (
	"github.com/ghgsnaidu/go-db-demo/pkg/controllers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
	r.HandleFunc("/users/", controllers.PostUser).Methods("POST")
}
