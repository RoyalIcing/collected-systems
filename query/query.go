package query

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"

	"github.com/RoyalIcing/collected-systems/types"
)

var schemaString = `
schema {
	query: Query
}

type Service {
	domain: String
}

interface Profile {
	service: Service
	username: String
}

type GitHubUserProfile implements Profile {
	service: Service
	username: String
	reposURL: ID
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

// PeopleArgs is the arguments take by a Person resolver
type PeopleArgs struct {
	RolesIn *[]string
}

// Resolver is the interface for concrete implementors
type Resolver interface {
	Hello() string
	People(ctx context.Context, args PeopleArgs) (*[]*types.Person, error)
	Services(ctx context.Context) (*[]*types.Service, error)
}

// MakeSchema creates a GraphQL schema from a query
func MakeSchema(resolver Resolver) *graphql.Schema {
	return graphql.MustParseSchema(schemaString, resolver)
}
