package app

import (
	"log"
	"micro-service-go/internal/config"
	"micro-service-go/internal/database"
	"micro-service-go/internal/delivery/http"
	"micro-service-go/internal/repositories"
	"micro-service-go/internal/services"
)

func Start() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db, err := database.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := http.NewUserHandler(userService)
	userRoute := http.NewUserRoute(userHandler)
	userRoute.Run(":" + cfg.HttpPort)
}
