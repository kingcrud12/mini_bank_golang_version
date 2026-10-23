package model

import (
	"learn_go_with_api/internal/domain/rule"
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func CreateCustomer(firstName string, lastName string, email string, password string) (*Customer, error) {
	if !rule.IsEmailValid(email) {
		return nil, ErrInvalidEmailFormat
	}
	return &Customer{
		ID:        uuid.NewString(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
	}, nil

}
