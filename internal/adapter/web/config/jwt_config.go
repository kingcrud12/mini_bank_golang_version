package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type JWTConfig struct {
	Secret     []byte
	Expiration time.Duration
	CookieName string
	Secure     bool
}

func LoadJWTConfig() JWTConfig {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading configuration from environment")
	}

	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET environment variable is required")
	}

	expirationMinutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRATION_MINUTES"))
	if err != nil || expirationMinutes <= 0 {
		expirationMinutes = 60
	}

	cookieName := os.Getenv("JWT_COOKIE_NAME")
	if cookieName == "" {
		cookieName = "access_token"
	}

	secure, _ := strconv.ParseBool(os.Getenv("JWT_COOKIE_SECURE"))

	return JWTConfig{
		Secret:     []byte(secret),
		Expiration: time.Duration(expirationMinutes) * time.Minute,
		CookieName: cookieName,
		Secure:     secure,
	}
}
