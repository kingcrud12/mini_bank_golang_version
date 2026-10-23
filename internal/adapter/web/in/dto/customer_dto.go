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
