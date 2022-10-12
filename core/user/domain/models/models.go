package models

import "time"

type User struct {
	ID        int64
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      string
	Photo     string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsDeleted bool
}

func (u *User) Validate() error {
	return nil
}
