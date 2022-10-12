package models

import (
	"github.com/solrac97gr/go-shop/core/product/domain/models"
	"time"
)

type Order struct {
	ID         int64
	PaymentID  int64
	BuyerID    int64
	OrderItems []*models.Product
	Amount     float64
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  time.Time
	IsDeleted  bool
}

func (o *Order) Validate() error {
	return nil
}
