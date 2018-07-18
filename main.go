package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/RoyalIcing/collected-systems/sources"
	"github.com/RoyalIcing/collected-systems/types"
)

type query struct{}

func (*query) Hello() string { return "Hello, world!" }

type peopleArgs struct {
	RolesIn *[]string
}

func (*query) People(ctx context.Context, args peopleArgs) (*[]*types.Person, error) {
	maybeAllPeople, err := sources.ReadPeopleCSVFile("./samples/cogent/people.csv")
	if err != nil {
		return nil, err
	}

	if maybeAllPeople == nil {
		return nil, nil
	}

	allPeople := *maybeAllPeople
	matchingPeople := allPeople[:]

	if allPeople != nil && args.RolesIn != nil {
		matchingPeople = allPeople[:0]
		for _, person := range allPeople {
			matchedRole := false
			for _, desiredRole := range *args.RolesIn {
				desiredRole = strings.ToUpper(desiredRole)
				if person.HasRole(desiredRole) {
					matchedRole = true
					break
				}
			}
			if matchedRole {
				matchingPeople = append(matchingPeople, person)
			}
		}
	}

	return &matchingPeople, nil
}

func (*query) Services(ctx context.Context) (*[]*types.Service, error) {
	maybeAllServices, err := sources.ReadServicesCSVFile("./samples/cogent/services.csv")
	if err != nil {
		return nil, err
	}

	if maybeAllServices == nil {
		return nil, nil
	}

	allServices := *maybeAllServices

	return &allServices, nil
}

func main() {
	s := `
	schema {
		query: Query
	}

	type Service {
		domain: String
	}

	type Profile {
		domain: String
		username: String
	}

	enum Role {
		EXECUTIVE
		ENGINEERING
		DESIGN
		PRODUCT
		PRINCIPAL
	}

	type Person {
		firstName: String
		lastName: String
		roles: [Role]
		profiles: [Profile]
	}

	type Query {
		hello: String!
		people(rolesIn: [Role!]): [Person]
		services: [Service]
	}
	`
	port := "3838"
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
