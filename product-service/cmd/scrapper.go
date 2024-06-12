package main

import (
	"database/sql"
	"log"
	"micro-service-go/internal/config"
	"micro-service-go/internal/database"
	"micro-service-go/internal/repositories"
	"micro-service-go/internal/services"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db, err := database.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)

	scrapping := services.NewScrappingHtml()

	outerHtml, err := scrapping.GetOuterHTML(
		"https://www.lazada.vn/catalog/?q=quan&page=3",
		[]string{
			"https://tpsservice-files-inner.cn-hangzhou.oss-cdn.aliyun-inc.com/images/resources/9dd6917e501f4144dd7af71009cceb63-1-1.png*",
			"https://tpsservice-files-inner.cn-hangzhou.oss-cdn.aliyun-inc.com/images/resources/9dd6917e501f4144dd7af71009cceb63-1-1.png*",
		},
	)
	if err != nil {
		log.Fatalf("Failed to get outer html: %v", err)
		return
	}

	list, err := scrapping.GetProductInfoFromList(outerHtml)
	if err != nil {
		return
	}

	err = productService.SaveProducts(list)
	if err != nil {
		log.Fatalf("Failed to save products: %v", err)
		return
	}
	// Add your scrapper logic here

}
