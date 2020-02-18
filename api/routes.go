package api

import (
	"dashboard/api/application"
	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	orderController application.OrderController
}

func (routes *Routes) Inject(orderController *application.OrderController) {
	routes.orderController = *orderController
}

func (routes *Routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/order", "order")
	registry.HandleGet("order", routes.orderController.Get)
}

