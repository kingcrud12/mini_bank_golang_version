package persistance

import (
	"time"

	"learn_go_with_api/internal/domain/model"
)

type CustomerEntity struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	FirstName string    `gorm:"column:first_name;not null"`
	LastName  string    `gorm:"column:last_name;not null"`
	Email     string    `gorm:"column:email;uniqueIndex;not null"`
	Password  string    `gorm:"column:password;not null"`
	CreatedAt time.Time `gorm:"column:created_at;not null"`
}

func (CustomerEntity) TableName() string {
	return "customers"
}

func ToCustomerEntity(customer *model.Customer) *CustomerEntity {
	return &CustomerEntity{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Password:  customer.Password,
		CreatedAt: customer.CreatedAt,
	}
}

func (e *CustomerEntity) ToDomain() *model.Customer {
	return &model.Customer{
		ID:        e.ID,
		FirstName: e.FirstName,
		LastName:  e.LastName,
		Email:     e.Email,
		Password:  e.Password,
		CreatedAt: e.CreatedAt,
	}
}
