package types

// Profile is an individual on a platform, e.g. GitHub, Medium
type Profile struct {
	domain   string
	username string
}

// NewProfile makes a person with the provided values
func NewProfile(domain string, username string) *Profile {
	user := Profile{
		domain:   domain,
		username: username,
	}
	return &user
}

// Service resolved
func (user *Profile) Service() *Service {
	return NewService(user.domain)
}

// Domain resolved
func (user *Profile) Domain() *string {
	return &user.domain
}

// Username resolved
func (user *Profile) Username() *string {
	return &user.username
}
