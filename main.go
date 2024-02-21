package main

import (
	"list/config"
	"list/controller"
	"list/database"
	"list/repository"
	"list/service"
	"log"
	"net/http"
)

func main() {

	cnf := config.Get()
	//redis connection
	databaseRedis := database.GetConnectionRedis(cnf)
	databaseDialRedis := database.GetConnectionDial(cnf)
	//postgres connection
	databasePostgres := database.GetConnectionPostgre(cnf)
	//repository
	cache := repository.NewCache(databaseRedis, databaseDialRedis)
	repo := repository.NewNoteRepository(databasePostgres)
	//service
	serv := service.NewCommonNoteService(cache, repo)
	//controller
	noteController := controller.NewNoteController(serv)

	noteController.RouteCommonNote()

	if err := http.ListenAndServe(cnf.Server.Port, nil); err != nil {
		log.Fatal(err)
	}
}
