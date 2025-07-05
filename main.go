package main

import (
	"log"
	"net/http"

	"assignment5/config"
	"assignment5/database"
	"assignment5/routes"
	"assignment5/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/viper"
)

func main() {
	// Load .env or system env
	config.LoadConfig()

	// Init Zap logger
	utils.InitLogger()
	defer utils.Logger.Sync()

	// Init PostgreSQL DB
	db := database.InitDB()
	defer db.Close()

	// Setup router with Chi
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Mount all routes under /api
	r.Mount("/", routes.SetupRoutes(db, utils.Logger))

	// Get PORT from env
	port := viper.GetString("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
