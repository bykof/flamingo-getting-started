package infrastructure

import (
	"dashboard/order/domain/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

type OrderService struct{}

func (orderService OrderService) GetOrder(id string) (entity.Order, error) {
	var order entity.Order
	response, err := http.Get(fmt.Sprintf("order-service/orders/%s", id))
	if err != nil {
		return order, err
	}
	defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(order)
	if err != nil {
		return order, err
	}
	return order, err
}
