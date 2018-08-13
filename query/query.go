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

interface Node {
	id: ID!
}


type Query {
	hello: String!
	people(rolesIn: [Role!]): [Person]
	services: [Service]

	examplePosts: PostsConnection!
}


enum MediaBaseType {
  TEXT
  IMAGE
  AUDIO
  VIDEO
  APPLICATION
}

# See https://www.iana.org/assignments/media-types/media-types.xhtml
type MediaType {
  baseType: String!
  subtype: String!
  parameters: [String]
}

interface Asset {
  mediaType: MediaType!
}

type AssetReference implements Node {
  id: ID!

  asset: Asset
}

type MarkdownDocument implements Asset {
  mediaType: MediaType!
  source: String

  #assetReferences: [AssetReference]
}



type PostsConnection {
  edges: [PostEdge]
}

type PostEdge {
  node: Post
  cursor: ID!
}

type Post implements Node {
  id: ID!

  content: MarkdownDocument
  #authors: [User]
  #title: String
  #createdAt: UTCTime
  #updatedAt: UTCTime
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
	ExamplePosts(ctx context.Context) (types.PostConnection, error)
}

// MakeSchema creates a GraphQL schema from a query
func MakeSchema(resolver Resolver) *graphql.Schema {
	return graphql.MustParseSchema(schemaString, resolver)
}
