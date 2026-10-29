package model

import "fmt"

type Money struct {
	currency Currency
	amount   int64
}

type Currency string

const (
	Euro   Currency = "Euro"
	Dollar Currency = "Dollar"
)

func (c Currency) isValid() bool {
	switch c {
	case Euro, Dollar:
		return true
	default:
		return false
	}
}

func NewMoney(newCurrency Currency, newAmount int64) (*Money, error) {
	if !newCurrency.isValid() {
		return nil, fmt.Errorf("invalid currency: %s", newCurrency)
	}
	if newAmount <= 0 {
		return nil, fmt.Errorf("amount must be positive, got %d", newAmount)
	}
	return &Money{
		currency: newCurrency,
		amount:   newAmount,
	}, nil

}
