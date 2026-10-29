package out

type TokenGenerator interface {
	Generate(customerID string, email string) (string, error)
	Validate(token string) (customerID string, err error)
}
