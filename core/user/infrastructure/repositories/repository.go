package repositories

import (
	"database/sql"
	"github.com/solrac97gr/go-shop/core/user/domain/models"
	"github.com/solrac97gr/go-shop/core/user/domain/ports"
)

type UserRepository struct {
	Database *sql.DB
}

var _ ports.UserRepository = &UserRepository{}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{Database: database}
}

func (u UserRepository) Create(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetByID(id string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetByEmail(email string) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) GetAll() ([]*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Update(user *models.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
