package models

type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Photos      []string
	CreatorID   string
	Stock       int64
}

func (p *Product) Validate() error {
	return nil
}
