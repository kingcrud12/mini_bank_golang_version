package service

import (
	"learn_go_with_api/internal/application/port/in"
	"learn_go_with_api/internal/application/port/out"
	"learn_go_with_api/internal/domain/model"
)

type GetProfileService struct {
	repo out.CustomerRepository
}

func NewGetProfileService(repo out.CustomerRepository) *GetProfileService {
	return &GetProfileService{repo: repo}
}

func (s *GetProfileService) GetProfile(customerID string) (*in.ProfileResult, error) {
	customer, err := s.repo.FindByID(customerID)
	if err != nil {
		return nil, err
	}

	if customer == nil {
		return nil, model.ErrCustomerNotFound
	}

	return &in.ProfileResult{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
	}, nil
}
