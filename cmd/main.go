package main

import (
	"GuardianPRO/pkg/handler"
	"GuardianPRO/pkg/repository"
	"GuardianPRO/pkg/service"
	"log"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	if err := runServer("8008", handlers); err != nil {
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
