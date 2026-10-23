package in

import (
	"learn_go_with_api/internal/adapter/web/in/dto"
	"learn_go_with_api/internal/domain/model"
)

type CustomerRegistered interface {
	Register(user dto.CustomerRegigestRequestDTO) (*model.Customer, error)
}
