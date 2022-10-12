package auth

type Credentials struct {
	Username string
	Password string
}

type Authenticator interface {
	Authenticate(credentials *Credentials) error
}

type CustomAuth struct {
}

var _ Authenticator = (*CustomAuth)(nil)

func NewCustomAuth() *CustomAuth {
	return &CustomAuth{}
}

func (c CustomAuth) Authenticate(credentials *Credentials) error {
	//TODO implement me
	panic("implement me")
}
