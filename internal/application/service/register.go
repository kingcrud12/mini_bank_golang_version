package service

import (
	"fmt"
	"learn_go_with_api/internal/adapter/web/in/dto"

	"learn_go_with_api/internal/application/port/out"
	"learn_go_with_api/internal/application/service/utils"
	"learn_go_with_api/internal/domain/model"
)

type RegisterService struct {
	repo           out.CustomerRepository
	passwordHasher out.PasswordHasher
}

func NewRegisterService(repo out.CustomerRepository, passwordHasher out.PasswordHasher) *RegisterService {
	return &RegisterService{repo: repo, passwordHasher: passwordHasher}
}

func (s *RegisterService) Register(user dto.CustomerRegigestRequestDTO) (*model.Customer, error) {

	if s.repo.ExistsByEmail(user.Email) {
		return nil, model.ErrEmailAlreadyUsed
	}

	passwordHasher := s.passwordHasher.Hash(user.Password)
	user.Password = passwordHasher

	userCreated, err := utils.BuildCustomerFromInput(user)

	if err != nil {
		return nil, fmt.Errorf("error during user creation: %w", err)
	}

	return s.repo.Save(userCreated)
}
