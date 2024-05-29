package repositories

import (
	"database/sql"
	"micro-service-go/internal/database"
	"micro-service-go/internal/models"
	"strconv"
)

type OrderRepository interface {
	GetOrders() ([]models.Order, error)
}

type orderRepository struct {
	DB *database.DB
}

func (u orderRepository) GetOrders() ([]models.Order, error) {
	query := `SELECT id, user_id, total FROM orders`
	rows, err := u.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)

	var orders []models.Order

	for rows.Next() {
		var order models.Order
		err := rows.Scan(&order.ID, &order.UserId, &order.Total)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	if len(orders) == 0 {
		return orders, nil
	}

	var str string
	for _, order := range orders {
		str += strconv.Itoa(order.ID) + ","
	}

	query = `SELECT id, order_id, product_id, quantity, price FROM order_items WHERE order_id IN (` + str[:len(str)-1] + `)`
	rows, err = u.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var orderItems []models.OrderItem
	for rows.Next() {
		var orderItem models.OrderItem
		err := rows.Scan(&orderItem.ID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price)
		if err != nil {
			return nil, err
		}
		orderItems = append(orderItems, orderItem)
	}

	orderItemMap := make(map[int][]models.OrderItem)
	for _, orderItem := range orderItems {
		orderItemMap[orderItem.OrderID] = append(orderItemMap[orderItem.OrderID], orderItem)
	}

	// Add the following code
	for i, order := range orders {
		orders[i].OrderItems = orderItemMap[order.ID]
	}

	return orders, nil
}

func NewOrderRepository(db *database.DB) OrderRepository {
	return &orderRepository{
		DB: db,
	}
}
