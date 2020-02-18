package infrastructure

import (
	"dashboard/order/domain/entity"
)

type FakeOrderService struct{}

func (fakeOrderService FakeOrderService) GetOrder(id string) (entity.Order, error) {
	return entity.Order{
		ID:     id,
		UserID: "1234",
		Sum:    12.4,
	}, nil
}
