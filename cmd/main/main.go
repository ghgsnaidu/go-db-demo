package main

import (
	"fmt"
	"net/http"

	"github.com/ghgsnaidu/go-db-demo/pkg/dbconfig"
	"github.com/ghgsnaidu/go-db-demo/pkg/routers"
	"github.com/ghgsnaidu/go-db-demo/pkg/services"
	"github.com/gorilla/mux"
)

func main() {

	//created router
	router := mux.NewRouter()
	//wire.Build(dbconfig.NewConfig, dbconfig.InitDB, services.InitServices)
	//getting db configurations

	fmt.Println("........router created")

	config := dbconfig.NewConfig()

	//opening db connc by injecting config and getting db instance
	fmt.Println("connecting to db")
	db := dbconfig.InitDB(config)

	//injecting db into services layer
	services.InitServices(db)

	routers.RegisterRoutes(router)
	http.Handle("/", router)

	fmt.Println("server starting......")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		fmt.Println("error starting server...")
		return
	}

}
