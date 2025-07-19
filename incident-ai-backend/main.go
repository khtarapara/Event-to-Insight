package main

import (
	"fmt"
	"incident-ai-backend/config"
	"incident-ai-backend/logger"
	"incident-ai-backend/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Incident Triage API
// @version 1.0
// @description Auto-triages and classifies incidents via OpenAI.
// @host localhost:5174
// @BasePath /
func main() {
	logger.InitLogger()
	db := config.InitDB()
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		logger.Logger.Error("Error loading .env file", err)
		os.Exit(1)
	}

	origin := os.Getenv("FRONTEND_ORIGIN")
	if origin == "" {
		origin = "http://localhost:5173"
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{origin},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.RegisterRoutes(r, db)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback default
	}

	r.Run(fmt.Sprintf(":%s", port))
}
