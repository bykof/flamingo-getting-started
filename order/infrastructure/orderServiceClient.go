package infrastructure

import (
	"encoding/json"
	"dashboard/order/domain"
	"fmt"
	"net/http"
)

type OrderServiceClient struct{}

func (orderServiceClient OrderServiceClient) GetOrder(id string) (domain.Order, error) {
	var err error
	var order domain.Order
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
