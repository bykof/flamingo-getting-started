package service

import (
	"dashboard/order/domain/entity"
)

type IOrderService interface {
	GetOrder(id string) (entity.Order, error)
}
