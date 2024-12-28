package main

import (
	"go-app/initializers"
	// "go-app/migrations"
	"go-app/controllers"
	_ "go-app/docs"
	"go-app/helpers"
	"go-app/repositories"
	"go-app/routes"
	"go-app/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	// "github.com/rs/zerolog/log"
	"os"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {

	initializers.LoadEnvs()
	db := initializers.DatabaseConnection()
	// migrations
	// migrations.MigrateModels(db)
	// log.Info().Msg("Started Server!")
	// Database
	validate := validator.New()

	// Repositories
	tagsRepository := repositories.NewTagsREpositoryImpl(db)

	// Services
	tagsService := services.NewTagsServiceImpl(tagsRepository, validate)

	// Controllers
	tagsController := controllers.NewTagsController(tagsService)

	// Routes
	routes := routes.NewRouter(tagsController)

	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: routes,
	}

	err := server.ListenAndServe()
	helpers.ErrorPanic(err)
}
