package main

import (
	"cafe-pos/config"
	"cafe-pos/internal/app/user"
	"cafe-pos/internal/infrastructure/database/postgres"
	entrepo "cafe-pos/internal/infrastructure/persistence/ent"
	fiberapp "cafe-pos/internal/interface/http/fiber"
	userhandler "cafe-pos/internal/interface/http/fiber/user"
	"log"

	"github.com/WhanSupakan/helpmez"
)

// Composition root
// This file will wire all dependencies together
// No business logic should be here

func main() {
	// TODO: Initialize configuration
	cfg := config.LoadConfig()
	helpmez.Print(cfg)
	// TODO: Initialize infrastructure (database, cache, messaging)
	db, err := postgres.NewEntClient(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := postgres.EntMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	defer db.Close()

	// Initialize repositories
	userRepo := entrepo.NewUserRepository(db)

	// Initialize usecases
	userUsecase := user.NewUserUsecase(userRepo)

	// Initialize handlers
	userHandler := userhandler.NewHandler(userUsecase)

	// Initialize HTTP delivery layer
	server := fiberapp.NewServer(cfg)
	fiberapp.NewRouter(server, userHandler)

	addr := ":" + cfg.Port
	log.Printf("Starting server on %s", addr)
	if err := server.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
