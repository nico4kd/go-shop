package config

type Configurator interface {
	SetEnvironmentVariables() error
	GetDatabaseConfig() DatabaseConfig
	GetPort() string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

type CustomConfigurator struct {
	Database    DatabaseConfig
	Environment string
	Port        string
}

func NewCustomConfigurator() *CustomConfigurator {
	return &CustomConfigurator{}
}

var _ Configurator = (*CustomConfigurator)(nil)

func (c *CustomConfigurator) GetDatabaseConfig() DatabaseConfig {
	//TODO implement me
	panic("implement me")
}

func (c *CustomConfigurator) SetEnvironmentVariables() error {
	//TODO implement me
	panic("implement me")
}

func (c *CustomConfigurator) GetPort() string {
	return c.Port
}
