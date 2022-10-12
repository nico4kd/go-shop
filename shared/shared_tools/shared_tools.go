package shared

import (
	"database/sql"
	"github.com/solrac97gr/go-shop/shared/config"
	"github.com/solrac97gr/go-shop/shared/logger"
	"github.com/solrac97gr/go-shop/shared/session"
	"github.com/solrac97gr/go-shop/shared/validator"
)

type Tools struct {
	Logger               logger.Logger
	Validator            validator.Validator
	EnvironmentVariables config.Configurator
	SessionManager       session.Session
	DatabaseClient       *sql.DB
}
