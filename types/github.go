package types

// GitHubUser is an individual on GitHub
type GitHubUser struct {
	name string
}

// NewGitHubUser makes a person with the provided values
func NewGitHubUser(name string) *GitHubUser {
	user := GitHubUser{
		name: name,
	}
	return &user
}

// Name resolved
func (user *GitHubUser) Name() *string {
	return &user.name
}
