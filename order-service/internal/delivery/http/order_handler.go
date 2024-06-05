package http

import (
	"micro-service-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewUserRoute(handle *OrderHandler) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.GET("/orders", handle.GetOrders)

	return route
}

type OrderHandler struct {
	orderService services.OrderService
}

func NewOrderHandler(orderService services.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (o *OrderHandler) GetOrders(c *gin.Context) {
	orders, err := o.orderService.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"orders": orders,
	})
}
