package model

import "errors"

var ErrEmailAlreadyUsed = errors.New("email already used")

var ErrInvalidEmailFormat = errors.New("invalid email format")

var ErrInvalidCredentials = errors.New("invalid email or password")

var ErrInvalidToken = errors.New("invalid or expired token")

var ErrCustomerNotFound = errors.New("customer not found")

var ErrHashPasswordFailed = errors.New("Erreur lors du hachge du mot de pase")

var ErrAccountAlreadyBlocked = errors.New("account is already blocked")
