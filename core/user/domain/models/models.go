package models

type User struct {
	ID        int64
	Email     string
	Password  string
	FirstName string
	LastName  string
	Role      string
	Photo     string
	Country   string
}

func (u *User) Validate() error {
	return nil
}
