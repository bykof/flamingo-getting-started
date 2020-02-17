package infrastructure

import "dashboard/order/domain"

type OrderServiceFakeClient struct{}

func (orderServiceFakeClient OrderServiceFakeClient) GetOrder(id string) (domain.Order, error) {
	return domain.Order{
		ID:     id,
		UserID: "1234",
		Sum:    12.4,
	}, nil
}
