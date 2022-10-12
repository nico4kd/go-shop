package application

import (
	"github.com/solrac97gr/go-shop/core/user/domain/models"
	"github.com/solrac97gr/go-shop/core/user/domain/ports"
)

type UserApplication struct {
	repository ports.UserRepository
}

var _ ports.UserApplication = (*UserApplication)(nil)

func NewUserApplication(repository ports.UserRepository) *UserApplication {
	return &UserApplication{repository: repository}
}

func (u UserApplication) Create(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserApplication) GetByID(id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserApplication) GetByEmail(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserApplication) GetAll() ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserApplication) Update(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserApplication) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
