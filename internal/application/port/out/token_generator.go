package out

type TokenGenerator interface {
	Generate(customerID string, email string) (string, error)
}
