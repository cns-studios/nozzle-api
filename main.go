package main

import (
	"context"
	"net/http"
	"nozzle-api/database"
	"nozzle-api/github-bot"
	"os"

	"nozzle-api/utils"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	utils.SetupLogging()
	utils.LogInfo("Starting application...")

	dbConnectionString := os.Getenv("DB_URL")
	pool, err := pgxpool.New(context.Background(), dbConnectionString)
	if err != nil {
		utils.LogError("Database connection failed: %v", err)
		return
	} else {
		utils.LogInfo("Database connection established")
	}
	defer pool.Close()
	database.RunMigration(context.Background(), pool)
	if os.Getenv("PORT") == "" {
		utils.LogError("PORT environment variable is not set")
		return
	}

	router := mux.NewRouter()

	router.HandleFunc("/api", utils.IsHealthy).Methods(http.MethodGet)
	router.HandleFunc("/api/repositories/{userid}", github_bot.GetAllrepos).Methods(http.MethodGet)

	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		utils.LogError("Server failed: %v", err)
	}
}
