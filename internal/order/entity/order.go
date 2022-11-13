package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	return &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}, nil
}

func (o *Order) isValid() error {
	if o.ID == "" {
		return errors.New("invalid id")
	}
	return nil
}
