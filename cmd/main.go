package main

import (
	"log"

	goweb "github.com/oleg5896/go-web"
	"github.com/oleg5896/go-web/pkg/handler"
	"github.com/oleg5896/go-web/pkg/repository"
	"github.com/oleg5896/go-web/pkg/service"
)

func main() {
	repos := repository.newRepository()
	services := service.newService(repos)
	handlers := handler.newHandler(services)

	srv := new(goweb.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		 log.Fatal("error start server")
	}
}
