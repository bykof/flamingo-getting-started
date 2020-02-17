package interfaces

import (
	"dashboard/order/domain"
)

type IOrderServiceClient interface {
	GetOrder(id string) (domain.Order, error)
}