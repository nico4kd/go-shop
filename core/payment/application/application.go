package application

import (
	"github.com/solrac97gr/go-shop/core/payment/domain/models"
	"github.com/solrac97gr/go-shop/core/payment/domain/ports"
)

type PaymentApplication struct {
	PaymentRepository ports.PaymentRepository
}

var _ ports.PaymentApplication = &PaymentApplication{}

func NewPaymentApplication(paymentRepository ports.PaymentRepository) *PaymentApplication {
	return &PaymentApplication{
		PaymentRepository: paymentRepository,
	}
}

func (p PaymentApplication) Create(payment *models.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentApplication) GetByID(id string) (*models.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentApplication) GetAll() ([]*models.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentApplication) Update(payment *models.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentApplication) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
