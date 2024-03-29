package main

import (
	"log"
	"shop/internal/handler"
	"shop/internal/repository"
	"shop/internal/routes"
	"shop/internal/service"
	"shop/pkg/database/postgres"

	"github.com/gofiber/fiber/v2"
)





func main(){
	app := fiber.New()
	pg, err := postgres.NewConnPostgres()
	if err != nil{
		log.Println(err)
	}
	repositories := repository.NewRepositories(pg)
	services := service.NewServices(repositories)
	handlers := handler.NewHandlers(*services)
	routes.InitRoutes(app, *handlers)
	err = app.Listen(":8080")
	if err != nil{
		log.Println(err)
	}
}