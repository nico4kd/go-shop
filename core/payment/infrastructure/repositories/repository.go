package repositories

import (
	"database/sql"
	"github.com/solrac97gr/go-shop/core/payment/domain/models"
	"github.com/solrac97gr/go-shop/core/payment/domain/ports"
)

type PaymentRepository struct {
	Database *sql.DB
}

var _ ports.PaymentRepository = &PaymentRepository{}

func NewPaymentRepository(database *sql.DB) *PaymentRepository {
	return &PaymentRepository{Database: database}
}

func (p PaymentRepository) Create(payment *models.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentRepository) GetByID(id string) (*models.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentRepository) GetAll() ([]*models.Payment, error) {
	//TODO implement me
	panic("implement me")
}

func (p PaymentRepository) Update(payment *models.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (p PaymentRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
