package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joaquincamara/cv-server/api"
	"github.com/joaquincamara/cv-server/internal/aboutMe"
	"github.com/joaquincamara/cv-server/internal/coolFeatures"
	"github.com/joaquincamara/cv-server/internal/devTechs"
	"github.com/joaquincamara/cv-server/internal/experience"
	"github.com/joaquincamara/cv-server/internal/personalProjects"
	"github.com/joaquincamara/cv-server/internal/postgres"
)

func main() {

	dbPool, _ := postgres.InitPoolConnection()
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = postgres.GetEnvs("SERVER_PORT")
	}

	aboutMeRepo := postgres.NewAboutMeRepository(dbPool)
	aboutMeService := aboutMe.NewAboutMeService(aboutMeRepo)
	aboutMeHandler := api.NewAboutMeHandler(aboutMeService)

	coolFeaturesRepo := postgres.NewCoolFeaturesRepository(dbPool)
	coolFeaturesService := coolFeatures.NewCoolFeaturesService(coolFeaturesRepo)
	coolFeaturesHandler := api.NewCoolFeaturesHandler(coolFeaturesService)

	devTechRepo := postgres.NewDevTechsRepository(dbPool)
	devTechService := devTechs.NewDevTechService(devTechRepo)
	devTechHandler := api.NewDevTechHandler(devTechService)

	experienceRepo := postgres.NewExperienceRepository(dbPool)
	experienceService := experience.NewExperienceService(experienceRepo)
	experienceHandler := api.NewExperienceHandler(experienceService)

	personalProjectsRepo := postgres.NewPersonalProjectsRepository(dbPool)
	personalProjectsService := personalProjects.NewCoolFeatureService(personalProjectsRepo)
	personalProjectsHandler := api.NewPersonalProjectHandler(personalProjectsService)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Post("/aboutMe", aboutMeHandler.Post)
	router.Get("/aboutMe", aboutMeHandler.GetAll)
	router.Delete("/aboutMe", aboutMeHandler.Delete)
	router.Put("/aboutMe", aboutMeHandler.Put)

	router.Post("/coolfeatures", coolFeaturesHandler.Post)
	router.Get("/coolfeatures", coolFeaturesHandler.GetAll)
	router.Delete("/coolfeatures", coolFeaturesHandler.Delete)
	router.Put("/coolfeatures", coolFeaturesHandler.Put)

	router.Post("/devTech", devTechHandler.Post)
	router.Get("/devTech", devTechHandler.GetAll)
	router.Delete("/devTech", devTechHandler.Delete)
	router.Put("/devTech", devTechHandler.Put)

	router.Post("/experience", experienceHandler.Post)
	router.Get("/experience", experienceHandler.GetAll)
	router.Delete("/experience", experienceHandler.Delete)
	router.Put("/experience", experienceHandler.Put)

	router.Post("/personalprojects", personalProjectsHandler.Post)
	router.Get("/personalprojects", personalProjectsHandler.GetAll)
	router.Delete("/personalprojects", personalProjectsHandler.Delete)
	router.Put("/personalprojects", personalProjectsHandler.Put)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
