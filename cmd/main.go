package main

import (
	"github.com/blazee5/EffectiveMobile_test/internal/app"
	"github.com/blazee5/EffectiveMobile_test/lib/logger"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("error while load .env: %v", err)
	}

	log := logger.NewLogger()

	app.Run(log)
}
