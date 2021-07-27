package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joaquincamara/cv-server/api"
	"github.com/joaquincamara/cv-server/internal/aboutMe"
	"github.com/joaquincamara/cv-server/internal/coolFeatures"
	"github.com/joaquincamara/cv-server/internal/devTechs"
	"github.com/joaquincamara/cv-server/internal/experience"
	"github.com/joaquincamara/cv-server/internal/personalProjects"
	"github.com/joaquincamara/cv-server/internal/postgres"
	"github.com/joaquincamara/silver"
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

	router := silver.NewRouter()

	router.POST("/aboutMe", aboutMeHandler.Post)
	router.GET("/aboutMe", aboutMeHandler.GetAll)
	router.DELETE("/aboutMe", aboutMeHandler.Delete)
	router.PUT("/aboutMe", aboutMeHandler.Put)

	router.POST("/coolfeatures", coolFeaturesHandler.Post)
	router.GET("/coolfeatures", coolFeaturesHandler.GetAll)
	router.DELETE("/coolfeatures", coolFeaturesHandler.Delete)
	router.PUT("/coolfeatures", coolFeaturesHandler.Put)

	router.POST("/devTech", devTechHandler.Post)
	router.GET("/devTech", devTechHandler.GetAll)
	router.DELETE("/devTech", devTechHandler.Delete)
	router.PUT("/devTech", devTechHandler.Put)

	router.POST("/experience", experienceHandler.Post)
	router.GET("/experience", experienceHandler.GetAll)
	router.DELETE("/experience", experienceHandler.Delete)
	router.PUT("/experience", experienceHandler.Put)

	router.POST("/personalprojects", personalProjectsHandler.Post)
	router.GET("/personalprojects", personalProjectsHandler.GetAll)
	router.DELETE("/personalprojects", personalProjectsHandler.Delete)
	router.PUT("/personalprojects", personalProjectsHandler.Put)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
