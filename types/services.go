package types

// Service is a platform used by the organization, e.g. GitHub, Medium
type Service struct {
	domain string
}

// NewService makes a service with the provided values
func NewService(domain string) *Service {
	user := Service{
		domain: domain,
	}
	return &user
}

// Domain resolved
func (user *Service) Domain() *string {
	return &user.domain
}
