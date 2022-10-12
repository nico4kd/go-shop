package factories

import (
	"github.com/solrac97gr/go-shop/shared/config"
	"github.com/solrac97gr/go-shop/shared/logger"
	"github.com/solrac97gr/go-shop/shared/session"
	"github.com/solrac97gr/go-shop/shared/shared_tools"
	"github.com/solrac97gr/go-shop/shared/validator"
)

type Factory interface {
	BuildSharedTools() *shared.Tools
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
	}
}
