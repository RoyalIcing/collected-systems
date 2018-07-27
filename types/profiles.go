package types

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// Profile is an individual on a platform, e.g. GitHub, Medium
type Profile struct {
	domain   string
	username string
}

// GitHubUserProfile is an account on GitHub
type GitHubUserProfile struct {
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

// Username resolved
func (user *Profile) Username() *string {
	return &user.username
}

// NewGitHubUserProfile makes a person with the provided values
func NewGitHubUserProfile(username string) *GitHubUserProfile {
	user := GitHubUserProfile{
		username: username,
	}
	return &user
}

// ToGitHubUserProfile attempts to convert to a GitHub user profile
func (user *Profile) ToGitHubUserProfile() (*GitHubUserProfile, bool) {
	if user.domain == "github.com" {
		return NewGitHubUserProfile(user.username), true
	}

	return nil, false
}

// Service resolved
func (user *GitHubUserProfile) Service() *Service {
	return NewService("github.com")
}

// Username resolved
func (user *GitHubUserProfile) Username() *string {
	return &user.username
}

// ReposURL resolved
func (user *GitHubUserProfile) ReposURL() *graphql.ID {
	id := graphql.ID("https://github.com/" + user.username + "?tab=repositories")
	return &id
}
