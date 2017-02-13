package campaignsvc

// Add database logic here

type Service struct{}

func (s Service) All() string {
	return "Get all campaigns"
}

func (s Service) One() string {
	return "Get one campaign"
}
