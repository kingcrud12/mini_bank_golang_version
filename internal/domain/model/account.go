package model

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            string
	CustomerId    string
	AccountNumber string
	Balance       Money
	AccountStatus AccountStatus
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func CreateNewAccount(customerId string, balance Money, accountNumber string) (*Account, error) {
	return &Account{
		ID:            uuid.NewString(),
		CustomerId:    customerId,
		AccountNumber: accountNumber,
		Balance:       balance,
		AccountStatus: AccountStatusActive,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}, nil

}

func (a *Account) Block() error {
	if a.AccountStatus == AccountStatusBlocked {
		return ErrAccountAlreadyBlocked
	}
	a.AccountStatus = AccountStatusBlocked
	a.UpdatedAt = time.Now().UTC()
	return nil
}
