package types

// Person is someone who has worked in your organization
type Person struct {
	firstName      string
	lastName       string
	roles          []*string
	gitHubUsername *string
}

// NewPerson makes a person with the provided values
func NewPerson(firstName string, lastName string, roles []string, gitHubUsername *string) *Person {
	person := Person{
		firstName:      firstName,
		lastName:       lastName,
		roles:          optionalStrings(roles),
		gitHubUsername: gitHubUsername,
	}
	return &person
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
	return &person.roles
}

// GitHubUser resolves using the `gitHubUsername` field
func (person *Person) GitHubUser() *GitHubUser {
	if person.gitHubUsername == nil {
		return nil
	}

	user := NewGitHubUser(*person.gitHubUsername)
	return user
}
