package main

import (
	"log"

	restgin "github.com/DiasOrazbaev/RestGIN"
	"github.com/DiasOrazbaev/RestGIN/pkg/handler"
	"github.com/DiasOrazbaev/RestGIN/pkg/repository"
	"github.com/DiasOrazbaev/RestGIN/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server := new(restgin.Server)
	if err := server.Run("3030", handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
