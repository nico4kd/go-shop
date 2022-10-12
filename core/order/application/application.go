package application

import (
	"github.com/solrac97gr/go-shop/core/order/domain/models"
	"github.com/solrac97gr/go-shop/core/order/domain/ports"
)

type OrderApplication struct {
	OrderRepository ports.OrderRepository
}

var _ ports.OrderApplication = &OrderApplication{}

func NewOrderApplication(orderRepository ports.OrderRepository) ports.OrderApplication {
	return &OrderApplication{OrderRepository: orderRepository}
}

func (o OrderApplication) Create(order *models.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderApplication) GetByID(id int64) (*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderApplication) GetAll() ([]*models.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (o OrderApplication) Update(order *models.Order) error {
	//TODO implement me
	panic("implement me")
}

func (o OrderApplication) Delete(order *models.Order) error {
	//TODO implement me
	panic("implement me")
}
