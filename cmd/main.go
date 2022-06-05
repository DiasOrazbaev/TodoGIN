package main

import (
	"log"

	restgin "github.com/DiasOrazbaev/RestGIN"
	"github.com/DiasOrazbaev/RestGIN/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)

	server := new(restgin.Server)
	if err := server.Run("3030", handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}
