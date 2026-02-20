package main

import (
	"log"

	"github.com/pleasure1234/my-gateway/config"
	"github.com/pleasure1234/my-gateway/router"
)

func main() {
	cfg := config.Load()

	r := router.SetupRouter()

	log.Printf("Server starting on port %s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}