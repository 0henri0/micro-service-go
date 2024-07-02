package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"micro-service-go/internal/config"
)

func Start() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	route.Run(":" + cfg.HttpPort)
}
