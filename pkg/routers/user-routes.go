package routers

import (
	user_controllers "github.com/ghgsnaidu/go-db-demo/pkg/rest-handlers"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users", user_controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users/{id}", user_controllers.GetUserById).Methods("GET")
	r.HandleFunc("/users/", user_controllers.PostUser).Methods("POST")
}
