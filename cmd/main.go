package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joaquincamara/cv-server/api"
	"github.com/joaquincamara/cv-server/internal/aboutMe"
	"github.com/joaquincamara/cv-server/internal/devTechs"
	"github.com/joaquincamara/cv-server/internal/postgres"
)

func main() {
	//
	dbPool, _ := postgres.InitPoolConnection()

	devTechRepo := postgres.NewDevTechsRepository(dbPool)
	devTechService := devTechs.NewDevTechService(devTechRepo)
	devTechHandler := api.NewDevTechHandler(devTechService)

	aboutMeRepo := postgres.NewAboutMeRepository(dbPool)
	aboutMeService := aboutMe.NewAboutMeService(aboutMeRepo)
	aboutMeHandler := api.NewAboutMeHandler(aboutMeService)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/devTech", devTechHandler.Post)
	router.Get("/devTech", devTechHandler.GetAll)
	router.Delete("/devTech", devTechHandler.Delete)
	router.Put("/devTech", devTechHandler.Put)

	router.Post("/aboutMe", aboutMeHandler.Post)
	router.Get("/aboutMe", aboutMeHandler.GetAll)
	router.Delete("/aboutMe", aboutMeHandler.Delete)
	router.Put("/aboutMe", aboutMeHandler.Put)

	log.Fatal(http.ListenAndServe(":8080", router))
}
