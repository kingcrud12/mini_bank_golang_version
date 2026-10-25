package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type BcryptPasswordHasher struct{}

func NewBcryptPasswordHasher() *BcryptPasswordHasher {
	return &BcryptPasswordHasher{}
}

func (h *BcryptPasswordHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Erreur lors du hachage du mot de mot de passe: %v", err)
	}
	return string(hashed), nil
}

func (h *BcryptPasswordHasher) Match(rawPassword string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(password), []byte(rawPassword)) == nil
}
