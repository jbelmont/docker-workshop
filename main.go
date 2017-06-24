package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbelmont/docker-workshop/model"
	"github.com/jbelmont/docker-workshop/redis"
	"github.com/jbelmont/docker-workshop/routes"
)

func getRouter() *mux.Router {
	return routes.NewRouter()
}

func initDB() {
	session := model.InitDB()
	model.CreateInitDocument(session)
}

func initRedis() {
	redis.ConnectRedis()
}

func main() {
	defer initDB()
	defer initRedis()
	router := getRouter()
	http.ListenAndServe(":3000", router)
}
