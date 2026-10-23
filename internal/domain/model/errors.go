package model

import "errors"

var ErrEmailAlreadyUsed = errors.New("email already used")

var ErrInvalidEmailFormat = errors.New("invalid email format")
