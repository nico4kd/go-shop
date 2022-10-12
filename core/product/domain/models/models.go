package models

import "time"

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Photos      []string
	CreatorID   string
	Stock       int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
	IsDeleted   bool
}

func (p *Product) Validate() error {
	return nil
}
