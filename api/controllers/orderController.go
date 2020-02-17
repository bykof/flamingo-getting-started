package controllers

import (
	"context"
	orderDomain "dashboard/order/domain"
	orderInterfaces "dashboard/order/interfaces"
	"flamingo.me/flamingo/v3/framework/web"
)

const orderIDKey = "orderID"

type (
	OrderController struct {
		responder          *web.Responder
		orderServiceClient orderInterfaces.IOrderServiceClient
	}

	GetOrderResponse struct {
		Order orderDomain.Order
	}
)

func (orderController *OrderController) Inject(
	responder *web.Responder,
	orderServiceClient orderInterfaces.IOrderServiceClient,
) {
	orderController.responder = responder
	orderController.orderServiceClient = orderServiceClient
}

func (orderController *OrderController) Get(ctx context.Context, req *web.Request) web.Result {
	orderID, err := req.Query1(orderIDKey)
	if err != nil {
		return orderController.responder.ServerError(err)
	}

	order, err := orderController.orderServiceClient.GetOrder(orderID)
	if err != nil {
		return orderController.responder.ServerError(err)
	}
	return orderController.responder.Data(GetOrderResponse{
		Order: order,
	})
}
