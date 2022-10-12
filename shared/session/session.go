package session

import "github.com/solrac97gr/go-shop/core/user/domain/models"

type Session interface {
	GetUser() models.User
	SetUser() models.User
	Expire() error
}

type Manager struct {
}

var _ Session = (*Manager)(nil)

func (m Manager) GetUser() models.User {
	//TODO implement me
	panic("implement me")
}

func (m Manager) SetUser() models.User {
	//TODO implement me
	panic("implement me")
}

func (m Manager) Expire() error {
	//TODO implement me
	panic("implement me")
}

func NewManager() *Manager {
	return &Manager{}
}
