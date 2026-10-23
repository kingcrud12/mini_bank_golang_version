package out

type PasswordHasher interface {
	Hash(password string) string
	Match(rawPassword string, password string) bool
}
