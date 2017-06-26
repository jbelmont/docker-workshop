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
	context := model.NewContext()
	collection := context.DBCollection()
	model.CreateInitDocument(collection)
}

func initRedis() {
	redis.ConnectRedis()
}

func init() {
	go initDB()
	go initRedis()
}

func main() {
	r := getRouter()
	http.ListenAndServe(":3000", r)
}
