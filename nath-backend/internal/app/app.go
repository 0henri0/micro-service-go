package app

import (
	"log"
	"nath-backend/internal/config"
	"nath-backend/internal/database"
	"nath-backend/internal/delivery/http"
	"nath-backend/internal/repositories"
	"nath-backend/internal/services"
)

func Start() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db, err := database.ConnectOrm(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := http.NewProductHandler(productService)
	route := http.NewProductRoute(productHandler)
	route.Run(":" + cfg.HttpPort)
}
