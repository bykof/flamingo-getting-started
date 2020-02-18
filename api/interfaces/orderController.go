package interfaces

import (
	"context"
	"dashboard/order/domain/entity"
	"dashboard/order/domain/service"
	"flamingo.me/flamingo/v3/framework/web"
)

const orderIDKey = "orderID"

type (
	OrderController struct {
		responder    *web.Responder
		orderService service.IOrderService
	}

	GetOrderResponse struct {
		Order entity.Order
	}
)

func (orderController *OrderController) Inject(
	responder *web.Responder,
	orderService service.IOrderService,
) {
	orderController.responder = responder
	orderController.orderService = orderService
}

func (orderController *OrderController) Get(ctx context.Context, req *web.Request) web.Result {
	orderID, err := req.Query1(orderIDKey)
	if err != nil {
		return orderController.responder.ServerError(err)
	}

	order, err := orderController.orderService.GetOrder(orderID)
	if err != nil {
		return orderController.responder.ServerError(err)
	}
	return orderController.responder.Data(GetOrderResponse{
		Order: order,
	})
}
