package logger

type Logger interface {
	Info(message string)
	Error(message string)
	Debug(message string)
	Warn(message string)
}

type CustomLogger struct {
}

func NewCustomLogger() *CustomLogger {
	return &CustomLogger{}
}

var _ Logger = (*CustomLogger)(nil)

func (c *CustomLogger) Info(message string) {
	//TODO implement me
	panic("implement me")
}

func (c *CustomLogger) Error(message string) {
	//TODO implement me
	panic("implement me")
}

func (c *CustomLogger) Debug(message string) {
	//TODO implement me
	panic("implement me")
}

func (c *CustomLogger) Warn(message string) {
	//TODO implement me
	panic("implement me")
}
