package types

// Person is someone who has worked in your organization
type Person struct {
	firstName      string
	lastName       string
	roles          []string
	gitHubUsername *string
}

// NewPerson makes a person with the provided values
func NewPerson(firstName string, lastName string, roles []string, gitHubUsername *string) *Person {
	person := Person{
		firstName:      firstName,
		lastName:       lastName,
		roles:          roles,
		gitHubUsername: gitHubUsername,
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

// GitHubUser resolves using the `gitHubUsername` field
func (person *Person) GitHubUser() *GitHubUser {
	if person.gitHubUsername == nil {
		return nil
	}

	user := NewGitHubUser(*person.gitHubUsername)
	return user
}
