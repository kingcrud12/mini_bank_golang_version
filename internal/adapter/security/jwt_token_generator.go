package security

import (
	"time"

	"learn_go_with_api/internal/adapter/web/config"
	"learn_go_with_api/internal/domain/model"

	"github.com/golang-jwt/jwt/v5"
)

type JWTTokenGenerator struct {
	config config.JWTConfig
}

func NewJWTTokenGenerator(cfg config.JWTConfig) *JWTTokenGenerator {
	return &JWTTokenGenerator{config: cfg}
}

func (g *JWTTokenGenerator) Generate(customerID string, email string) (string, error) {
	now := time.Now().UTC()

	claims := jwt.RegisteredClaims{
		Subject:   customerID,
		ID:        email,
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(g.config.Expiration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(g.config.Secret)
}

func (g *JWTTokenGenerator) Validate(tokenString string) (string, error) {
	claims := &jwt.RegisteredClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, model.ErrInvalidToken
		}
		return g.config.Secret, nil
	})

	if err != nil || !token.Valid {
		return "", model.ErrInvalidToken
	}

	return claims.Subject, nil
}
