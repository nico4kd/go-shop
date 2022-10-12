package factories

import (
	"database/sql"
	orderApp "github.com/solrac97gr/go-shop/core/order/application"
	orderPorts "github.com/solrac97gr/go-shop/core/order/domain/ports"
	orderHdl "github.com/solrac97gr/go-shop/core/order/infrastructure/handlers"
	orderRepo "github.com/solrac97gr/go-shop/core/order/infrastructure/repositories"
	paymentApp "github.com/solrac97gr/go-shop/core/payment/application"
	paymentPorts "github.com/solrac97gr/go-shop/core/payment/domain/ports"
	paymentHdl "github.com/solrac97gr/go-shop/core/payment/infrastructure/handlers"
	paymentRepo "github.com/solrac97gr/go-shop/core/payment/infrastructure/repositories"
	productApp "github.com/solrac97gr/go-shop/core/product/application"
	productPorts "github.com/solrac97gr/go-shop/core/product/domain/ports"
	productHdl "github.com/solrac97gr/go-shop/core/product/infrastructure/handlers"
	productRepo "github.com/solrac97gr/go-shop/core/product/infrastructure/repositories"
	userApp "github.com/solrac97gr/go-shop/core/user/application"
	userPorts "github.com/solrac97gr/go-shop/core/user/domain/ports"
	userHdl "github.com/solrac97gr/go-shop/core/user/infrastructure/handlers"
	userRepo "github.com/solrac97gr/go-shop/core/user/infrastructure/repositories"
	"github.com/solrac97gr/go-shop/shared/auth"

	"github.com/solrac97gr/go-shop/shared/config"
	"github.com/solrac97gr/go-shop/shared/logger"
	"github.com/solrac97gr/go-shop/shared/middleware"
	"github.com/solrac97gr/go-shop/shared/session"
	"github.com/solrac97gr/go-shop/shared/shared_tools"
	"github.com/solrac97gr/go-shop/shared/validator"
)

type Factory interface {
	BuildSharedTools() *shared.Tools
	BuildMiddlewares(tools *shared.Tools) middleware.Middleware
	BuildUserHandlers(tools *shared.Tools) userPorts.UserHandler
	BuildProductHandlers(tools *shared.Tools) productPorts.ProductHandler
	BuildPaymentHandlers(tools *shared.Tools) paymentPorts.PaymentHandler
	BuildOrderHandlers(tools *shared.Tools) orderPorts.OrderHandler
}

type CustomFactory struct {
}

var _ Factory = (*CustomFactory)(nil)

func NewCustomFactory() *CustomFactory {
	return &CustomFactory{}
}
func (f *CustomFactory) BuildSharedTools() *shared.Tools {
	return &shared.Tools{
		Logger:               logger.NewCustomLogger(),
		Validator:            validator.NewCustomValidator(),
		EnvironmentVariables: config.NewCustomConfigurator(),
		SessionManager:       session.NewManager(),
		DatabaseClient:       &sql.DB{},
	}
}

func (f *CustomFactory) BuildMiddlewares(tools *shared.Tools) middleware.Middleware {
	authenticator := auth.NewCustomAuth()
	return middleware.NewFiberMiddleware(authenticator)
}

func (f *CustomFactory) BuildUserHandlers(tools *shared.Tools) userPorts.UserHandler {
	repo := userRepo.NewUserRepository(tools.DatabaseClient)
	app := userApp.NewUserApplication(repo)
	return userHdl.NewUserHandlers(app)
}

func (f *CustomFactory) BuildProductHandlers(tools *shared.Tools) productPorts.ProductHandler {
	repo := productRepo.NewProductRepository(tools.DatabaseClient)
	app := productApp.NewProductApplication(repo)
	return productHdl.NewProductHandlers(app)
}

func (f *CustomFactory) BuildPaymentHandlers(tools *shared.Tools) paymentPorts.PaymentHandler {
	repo := paymentRepo.NewPaymentRepository(tools.DatabaseClient)
	app := paymentApp.NewPaymentApplication(repo)
	return paymentHdl.NewPaymentHandlers(app)
}

func (f *CustomFactory) BuildOrderHandlers(tools *shared.Tools) orderPorts.OrderHandler {
	repo := orderRepo.NewOrderRepository(tools.DatabaseClient)
	app := orderApp.NewOrderApplication(repo)
	return orderHdl.NewOrderHandlers(app)
}
