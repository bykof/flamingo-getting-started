package api

import (
	"dashboard/order/domain/service"
	"dashboard/order/infrastructure"
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"os"
)

type Module struct{}

func (module *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(Routes))
	if os.Getenv("fake") == "true" {
		injector.Bind(new(service.IOrderService)).To(infrastructure.FakeOrderService{})
	} else {
		injector.Bind(new(service.IOrderService)).To(infrastructure.OrderService{})
	}
}
