package api

import (
	"dashboard/api/controllers"
	"flamingo.me/flamingo/v3/framework/web"
)

type Routes struct {
	orderController controllers.OrderController
}

func (routes *Routes) Inject(orderController *controllers.OrderController) {
	routes.orderController = *orderController
}

func (routes *Routes) Routes(registry *web.RouterRegistry) {
	registry.Route("/order", "order")
	registry.HandleGet("order", routes.orderController.Get)
}

