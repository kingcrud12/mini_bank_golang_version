package out

import "learn_go_with_api/internal/domain/model"

type CustomerRepository interface {
	Save(*model.Customer) (*model.Customer, error)
	FindByEmail(email string) (*model.Customer, error)
	FindByID(id string) (*model.Customer, error)
	ExistsByEmail(email string) bool
}
