package http

import (
	"nath-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewProductRoute(handle *ProductHandler) *gin.Engine {
	route := gin.Default()
	route.GET("/products/:id", ValidateGetProduct(), handle.GetProduct)

	return route
}

type ProductHandler struct {
	productService services.ProductService
}

func NewProductHandler(productService services.ProductService) *ProductHandler {
	return &ProductHandler{productService: productService}
}

func ValidateGetProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		_, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
			c.Abort()
		}
		c.Next()
	}
}

func (u *ProductHandler) GetProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	product, err := u.productService.Find(idInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"product": product,
	})
}
