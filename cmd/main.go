package main

import (
	"github.com/blazee5/EffectiveMobile_test/internal/app"
	"github.com/blazee5/EffectiveMobile_test/lib/logger"
	"github.com/joho/godotenv"
	"log"
)

// @title EffectiveMobile Test API
// @description This is a sample server for a car service.
// @version 1.0
// @host localhost:3000
// @BasePath /api/v1
func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error while load .env: %v", err)
	}

	log := logger.NewLogger()

	app.Run(log)
}
