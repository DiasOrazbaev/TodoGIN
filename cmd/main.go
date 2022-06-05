package main

import (
	"log"

	restgin "github.com/DiasOrazbaev/RestGIN"
	"github.com/DiasOrazbaev/RestGIN/pkg/handler"
	"github.com/DiasOrazbaev/RestGIN/pkg/repository"
	"github.com/DiasOrazbaev/RestGIN/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Failed init config: %s\n", err.Error())
	}

	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	server := new(restgin.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalln(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
