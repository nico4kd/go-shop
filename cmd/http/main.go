package main

import (
	"github.com/solrac97gr/go-shop/shared/factories"
	"github.com/solrac97gr/go-shop/shared/server"
)

func main() {
	factory := factories.NewCustomFactory()

	tools := factory.BuildSharedTools()
	middlewares := factory.BuildMiddlewares(tools)
	userHdl := factory.BuildUserHandlers(tools)
	productHdl := factory.BuildProductHandlers(tools)
	paymentHdl := factory.BuildPaymentHandlers(tools)
	orderHdl := factory.BuildOrderHandlers(tools)

	fiberServer := server.NewFiberServer(
		userHdl,
		productHdl,
		paymentHdl,
		orderHdl,
		middlewares,
	)

	fiberServer.Start(tools.EnvironmentVariables.GetPort())
}
