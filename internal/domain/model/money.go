package model

import "fmt"

type Money struct {
	currency string
	amount   int64
}

type curreny string

Const(
	Euro 
)

func NewMoney(newCurrency string, newAmount int64) (*Money, error) {
	if newAmount <= 0 {
		return nil, fmt.Errorf("amount must be positive, got %d", newAmount)
	}
	return &Money{
		currency: newCurrency,
		amount:   newAmount,
	}, nil

}
