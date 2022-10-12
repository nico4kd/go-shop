package repositories

import (
	"database/sql"
	"github.com/solrac97gr/go-shop/core/order/domain/models"
	"github.com/solrac97gr/go-shop/core/order/domain/ports"
)

type OrderRepository struct {
	Database *sql.DB
}

var _ ports.OrderRepository = &OrderRepository{}

func NewOrderRepository(database *sql.DB) *OrderRepository {
	return &OrderRepository{Database: database}
}

func (o OrderRepository) Create(order *models.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) GetByID(id int64) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) GetAll() ([]*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) Update(order *models.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderRepository) Delete(order *models.Order) error {
	//TODO implement me
	panic("implement me")
}
