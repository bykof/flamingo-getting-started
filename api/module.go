package api

import (
	orderInfrastructure "dashboard/order/infrastructure"
	orderInterfaces "dashboard/order/interfaces"
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"os"
)

type Module struct{}

func (module *Module) Configure(injector *dingo.Injector) {
	web.BindRoutes(injector, new(Routes))
	var orderServiceClient orderInterfaces.IOrderServiceClient
	orderServiceClient = orderInfrastructure.OrderServiceClient{}

	if os.Getenv("fake") == "true" {
		orderServiceClient = orderInfrastructure.OrderServiceFakeClient{}
	}
	injector.Bind(new(orderInterfaces.IOrderServiceClient)).To(orderServiceClient)
}
