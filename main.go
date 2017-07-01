package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbelmont/docker-workshop/handlers"
	"github.com/jbelmont/docker-workshop/model"
	"github.com/jbelmont/docker-workshop/redis"
	"github.com/jbelmont/docker-workshop/routes"
)

func getRouter() *mux.Router {
	return routes.NewRouter()
}

func main() {
	model.NewContext()
	redis.ConnectRedis()
	router := getRouter()
	router.HandleFunc("/ap/v1/users", handlers.GetUsers)
	log.Fatal(http.ListenAndServe(":3001", router))
}
