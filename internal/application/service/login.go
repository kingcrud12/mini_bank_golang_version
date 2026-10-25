package service

import (
	"learn_go_with_api/internal/adapter/web/in/dto"
	"learn_go_with_api/internal/application/port/in"
	"learn_go_with_api/internal/application/port/out"
	"learn_go_with_api/internal/domain/model"
)

type LoginService struct {
	repo           out.CustomerRepository
	passwordHasher out.PasswordHasher
	tokenGenerator out.TokenGenerator
}

func NewLoginService(repo out.CustomerRepository, passwordHasher out.PasswordHasher, tokenGenerator out.TokenGenerator) *LoginService {
	return &LoginService{repo: repo, passwordHasher: passwordHasher, tokenGenerator: tokenGenerator}
}

func (s *LoginService) Login(credentials dto.CustomerLoginRequestDTO) (*in.LoginResult, error) {
	customer, err := s.repo.FindByEmail(credentials.Email)
	if err != nil {
		return nil, err
	}

	if customer == nil || !s.passwordHasher.Match(credentials.Password, customer.Password) {
		return nil, model.ErrInvalidCredentials
	}

	token, err := s.tokenGenerator.Generate(customer.ID, customer.Email)
	if err != nil {
		return nil, err
	}

	return &in.LoginResult{
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Token:     token,
	}, nil
}
