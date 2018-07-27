package sources

import (
	"context"

	"github.com/gobuffalo/packr"

	"github.com/RoyalIcing/collected-systems/query"
	"github.com/RoyalIcing/collected-systems/types"
)

// LocalFileSource reads from the file system
type LocalFileSource struct {
	box packr.Box
}

// NewLocalFileSource makes a new source from the local file system
func NewLocalFileSource(box packr.Box) LocalFileSource {
	return LocalFileSource{
		box: box,
	}
}

// Hello resolved
func (l LocalFileSource) Hello() string { return "Hello, world!" }

// People resolved
func (l LocalFileSource) People(ctx context.Context, args query.PeopleArgs) (*[]*types.Person, error) {
	file, err := l.box.Open("./people.csv")
	if err != nil {
		return nil, err
	}

	maybeAllPeople, err := ReadPeopleCSVFrom(file)
	if err != nil {
		return nil, err
	}

	if maybeAllPeople == nil {
		return nil, nil
	}

	return query.FilterPeople(*maybeAllPeople, args)
}

// Services resolved
func (l LocalFileSource) Services(ctx context.Context) (*[]*types.Service, error) {
	file, err := l.box.Open("./services.csv")
	if err != nil {
		return nil, err
	}

	maybeAllServices, err := ReadServicesCSVFrom(file)
	if err != nil {
		return nil, err
	}

	if maybeAllServices == nil {
		return nil, nil
	}

	allServices := *maybeAllServices

	return &allServices, nil
}
