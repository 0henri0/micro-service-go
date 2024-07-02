package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"micro-service-go/internal/config"
	"micro-service-go/internal/database"
	"micro-service-go/internal/models"
	"micro-service-go/internal/repositories"
	"micro-service-go/internal/services"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db := database.ConnectOrm(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *gorm.DB) {
		dbSql, _ := db.DB()
		dbSql.Close()
		if err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
	}(db)

	productRepository := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepository)
	scrapping := services.NewScrappingHtml()
	file, _ := os.Open("text.txt")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	wg := &sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		url := `https://www.lazada.vn/catalog/?q=quan&page=` + strconv.Itoa(i)
		scrappingHtml(url, scrapping, productService, wg)
	}

	wg.Wait()
}

func scrappingHtml(url string, scrapping services.ScrappingHtml, productService services.ProductService, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Scraping url: ", url)
	count := 0
	for {
		time.Sleep(time.Second * 2)
		var list []models.Product
		count++
		outerHtml, err := scrapping.GetOuterHTML(
			url,
			[]string{
				"https://tpsservice-files-inner.cn-hangzhou.oss-cdn.aliyun-inc.com/images/resources/9dd6917e501f4144dd7af71009cceb63-1-1.png*",
				"https://tpsservice-files-inner.cn-hangzhou.oss-cdn.aliyun-inc.com/images/resources/9dd6917e501f4144dd7af71009cceb63-1-1.png*",
			},
		)
		if err != nil {
			log.Fatalf("Failed to get outer html: %v", err)
			return
		}
		list, err = scrapping.GetProductInfoFromList(outerHtml)
		fmt.Println(list)
		if err != nil {
			return
		}
		err = productService.SaveProducts(list)
		if err != nil {
			log.Fatalf("Failed to save products: %v", err)
			return
		}

		if len(list) != 0 || count > 5 {
			break
		}
	}
}
