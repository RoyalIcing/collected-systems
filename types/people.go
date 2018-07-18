package types

// Person is someone who has worked in your organization
type Person struct {
	firstName        string
	lastName         string
	roles            []string
	profileUsernames map[string]string
}

// NewPerson makes a person with the provided values
func NewPerson(firstName string, lastName string, roles []string, profileUsernames map[string]string) *Person {
	person := Person{
		firstName:        firstName,
		lastName:         lastName,
		roles:            roles,
		profileUsernames: profileUsernames,
	}
	return &person
}

// HasRole returns true if has passed role
func (person *Person) HasRole(role string) bool {
	for _, r := range person.roles {
		if role == r {
			return true
		}
	}
	return false
}

// FirstName resolved
func (person *Person) FirstName() *string {
	return &person.firstName
}

// LastName resolved
func (person *Person) LastName() *string {
	return &person.lastName
}

// Roles resolved
func (person *Person) Roles() *[]*string {
	roles := optionalStrings(person.roles)
	return &roles
}

// Profiles resolves using the `profileUsernames` field
func (person *Person) Profiles() *[]*Profile {
	profiles := []*Profile{}
	for domain, username := range person.profileUsernames {
		profiles = append(profiles, NewProfile(domain, username))
	}
	return &profiles
}

// GitHubUser resolves using the `profileUsernames` field for `github.com`
func (person *Person) GitHubUser() *Profile {
	gitHubUsername, ok := person.profileUsernames["github.com"]
	if !ok {
		return nil
	}

	user := NewProfile("github.com", gitHubUsername)
	return user
}
