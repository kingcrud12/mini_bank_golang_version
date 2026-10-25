package out

type PasswordHasher interface {
	Hash(password string) (string, error)
	Match(rawPassword string, password string) bool
}
