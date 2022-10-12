package models

import "time"

type Payment struct {
	ID        string
	Amount    float64
	Status    string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeletedAt time.Time
	OrderID   int64
	IsDeleted bool
}

func (p *Payment) Validate() error {
	return nil
}
