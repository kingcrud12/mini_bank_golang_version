package dto

type CustomerRegigestRequestDTO struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type CustomerRegisterResponseDTO struct {
	FirstName string
	LastName  string
	Email     string
}

type CustomerLoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CustomerLoginResponseDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
