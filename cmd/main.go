package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joaquincamara/cv-server/api"
	"github.com/joaquincamara/cv-server/internal/devTechs"
	"github.com/joaquincamara/cv-server/internal/postgres"
)

func main() {

	dbPool, _ := postgres.InitPoolConnection()

	devTechRepo := postgres.NewDevTechsRepository(dbPool)
	devTechService := devTechs.NewDevTechService(devTechRepo)
	devTechHandler := api.NewDevTechHandler(devTechService)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/devTech", devTechHandler.Post)
	router.Get("/devTech", devTechHandler.GetAll)
	router.Delete("/devTech", devTechHandler.Delete)
	router.Put("/devTech", devTechHandler.Put)

	log.Fatal(http.ListenAndServe(":8080", router))
}
