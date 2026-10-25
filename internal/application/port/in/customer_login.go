package in

import "learn_go_with_api/internal/adapter/web/in/dto"

type CustomerLogin interface {
	Login(credentials dto.CustomerLoginRequestDTO) (*LoginResult, error)
}

type LoginResult struct {
	FirstName string
	LastName  string
	Email     string
	Token     string
}
