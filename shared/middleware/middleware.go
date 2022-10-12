package middleware

type Middleware interface {
}

type FiberMiddleware struct {
}

func NewFiberMiddleware() *FiberMiddleware {
	return &FiberMiddleware{}
}
