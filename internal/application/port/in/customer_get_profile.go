package in

type CustomerGetProfile interface {
	GetProfile(customerID string) (*ProfileResult, error)
}

type ProfileResult struct {
	FirstName string
	LastName  string
	Email     string
}
