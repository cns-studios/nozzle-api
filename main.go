package main

import (
	"context"
	"nozzle-api/database"
	"os"

	"nozzle-api/utils"

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
	}
	defer pool.Close()

	database.RunMigration(context.Background(), pool)
}
