package api

import (
	"dashboard/api/interfaces"
	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	orderController interfaces.OrderController
}

func (routes *Routes) Inject(orderController *interfaces.OrderController) {
	routes.orderController = *orderController
}

func (routes *Routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/order", "order")
	registry.HandleGet("order", routes.orderController.Get)
}

