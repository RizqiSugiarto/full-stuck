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
	databaseRedis := database.GetConnectionRedis(cnf)
	databaseDialRedis := database.GetConnectionDial(cnf)
	repo := repository.NewCache(databaseRedis, databaseDialRedis)
	serv := service.NewCommonNoteService(repo)
	noteController := controller.NewNoteController(serv)

	noteController.RouteCommonNote()

	if err := http.ListenAndServe(cnf.Server.Port, nil); err != nil {
		log.Fatal(err)
	}
}
