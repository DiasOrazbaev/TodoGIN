package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"os"

	restgin "github.com/DiasOrazbaev/RestGIN"
	"github.com/DiasOrazbaev/RestGIN/pkg/handler"
	"github.com/DiasOrazbaev/RestGIN/pkg/repository"
	"github.com/DiasOrazbaev/RestGIN/pkg/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Failed init config: %s\n", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Failed to loading env variables: %s\n", err.Error())
	}

	logrus.Println("Config correct initialized")

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("Failed to connection to db: %s\n", err.Error())
	}

	logrus.Println("Correctly connected to DB")

	repos := repository.NewRepository(db)
	newService := service.NewService(repos)
	handlers := handler.NewHandler(newService)

	server := new(restgin.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalln(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
