package main

import (
	"cafe-pos/config"
	"cafe-pos/internal/app/user"
	"cafe-pos/internal/infrastructure/database/postgres"
	"cafe-pos/internal/infrastructure/logger/std"
	entrepo "cafe-pos/internal/infrastructure/persistence/ent"
	fiberapp "cafe-pos/internal/interface/http/fiber"
	fibererrors "cafe-pos/internal/interface/http/fiber/errors"
	userhandler "cafe-pos/internal/interface/http/fiber/user"
	"log"

	"github.com/WhanSupakan/helpmez"
)

func main() {
	cfg := config.LoadConfig()
	helpmez.Print(cfg)

	db, err := postgres.NewEntClient(cfg.DatabaseDSN)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := postgres.EntMigrate(db); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	defer db.Close()

	logger := std.NewStdLogger()
	fibererrors.SetDefaultHandler(logger)

	userRepo := entrepo.NewUserRepository(db)
	userUsecase := user.NewUserUsecase(userRepo)
	userHandler := userhandler.NewHandler(userUsecase)

	server := fiberapp.NewServer(cfg)
	fiberapp.NewRouter(server, userHandler)

	addr := ":" + cfg.Port
	log.Printf("Starting server on %s", addr)
	if err := server.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
