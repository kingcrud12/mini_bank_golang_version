package utils

import (
	"learn_go_with_api/internal/adapter/web/in/dto"
	"learn_go_with_api/internal/domain/model"
)

func BuildCustomerFromInput(user dto.CustomerRegigestRequestDTO) (*model.Customer, error) {
	return model.CreateCustomer(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
	)

}
