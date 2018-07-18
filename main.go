package main

import (
	"context"
	// "encoding/csv"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"

	"./sources"
	"./types"
)

type query struct{}

func (*query) Hello() string { return "Hello, world!" }

func (*query) People(context.Context) (*[]*types.Person, error) {
	return sources.ReadPeopleCSVFile("./samples/cogent/people.csv")
}

func main() {
	s := `
	schema {
		query: Query
	}

	type GitHubUser {
		name: String
	}

	type Person {
		firstName: String
		lastName: String
		roles: [String]
		gitHubUser: GitHubUser
	}

	type Query {
		hello: String!
		people: [Person]
	}
	`
	port := "3838"
	schema := graphql.MustParseSchema(s, &query{})
	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
