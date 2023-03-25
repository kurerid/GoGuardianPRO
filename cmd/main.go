package main

import (
	"GuardianPRO/pkg/handler"
	"GuardianPRO/pkg/repository"
	"GuardianPRO/pkg/service"
	"fmt"
	"log"
)

func main() {
	cfg := repository.Config{
		ServerConfig: repository.ServerConfig{Port: "8080"},
		DbConfig: repository.DbConfig{
			Host:     "localhost",
			Port:     "5432",
			Username: "postgres",
			Password: "1234",
			DbName:   "GuardianPRO",
			SSLMode:  "disable",
		},
	}
	db, err := repository.NewPostgresDb(&cfg)
	if err != nil {
		fmt.Println("Сломались на подключении")
		log.Fatalln(err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	if err := runServer(cfg.ServerConfig.Port, handlers); err != nil {
		log.Fatalln(err)
	}
}

func runServer(port string, handler *handler.Handler) error {
	router := handler.InitRoutes()
	if err := router.Run(":" + port); err != nil {
		return err
	}
	return nil
}
