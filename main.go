package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jbelmont/docker-workshop/model"
	"github.com/jbelmont/docker-workshop/redis"
	"github.com/jbelmont/docker-workshop/routes"
)

func getRouter() *mux.Router {
	return routes.NewRouter()
}

func initRedis() {
	users := model.GetModels()
	for key, user := range users {
		setName := "user:" + strconv.Itoa(key)
		redis.SetHash(setName, user)
	}
}

func main() {
	initRedis()
	context := model.MongoSetup()
	defer context.Close()
	router := getRouter()
	log.Fatal(http.ListenAndServe(":3001", router))
}
