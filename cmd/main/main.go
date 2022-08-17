package main

import (
	"fmt"
	"net/http"

	"github.com/ghgsnaidu/go-db-demo/pkg/dbconfig"
	"github.com/ghgsnaidu/go-db-demo/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	fmt.Println("connecting to db")
	dbconfig.Init()
	routes.RegisterRoutes(router)
	http.Handle("/", router)

	fmt.Println("server starting......")
	err := http.ListenAndServe("localhost:3000", router)
	if err != nil {
		fmt.Println("error starting server...")
		return
	}

}
