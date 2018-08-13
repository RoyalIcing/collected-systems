package types

import (
	graphql "github.com/graph-gophers/graphql-go"
)

// MarkdownDocument is a text/markdown document
type MarkdownDocument struct {
	source string
}

// NewMarkdownDocument makes a Markdown
func NewMarkdownDocument(source string) *MarkdownDocument {
	markdownDocument := MarkdownDocument{
		source: source,
	}
	return &markdownDocument
}

// MediaType resolved
func (markdownDocument *MarkdownDocument) MediaType() MediaType {
	parameters := []string{}
	mediaType := NewMediaType("text", "markdown", parameters)
	return mediaType
}

// Source resolved
func (markdownDocument *MarkdownDocument) Source() *string {
	return &markdownDocument.source
}

// Post has a markdown document
type Post struct {
	id      string
	content *MarkdownDocument
}

// NewPost makes a post with the provided values
func NewPost(id string, content *MarkdownDocument) *Post {
	post := Post{
		id:      id,
		content: content,
	}
	return &post
}

// Title resolved
func (post *Post) Title() *string {
	title := "Example"
	return &title
}

// ID resolved
func (post *Post) ID() graphql.ID {
	return graphql.ID(post.id)
}

// Content resolved
func (post *Post) Content() *MarkdownDocument {
	return post.content
}

// PostEdge is a reference to a post within a connection
type PostEdge struct {
	post   *Post
	cursor string
}

// NewPostEdge makes a post edge with the provided values
func NewPostEdge(post *Post, cursor string) *PostEdge {
	postEdge := PostEdge{
		post:   post,
		cursor: cursor,
	}
	return &postEdge
}

// Node resolved
func (postEdge *PostEdge) Node() *Post {
	return postEdge.post
}

// Cursor resolved
func (postEdge *PostEdge) Cursor() graphql.ID {
	return graphql.ID(postEdge.cursor)
}

// PostConnection is a connection to a collection of posts
type PostConnection struct {
	edges *[]*PostEdge
}

// NewPostConnection makes a post connection with the provided values
func NewPostConnection(edges *[]*PostEdge) PostConnection {
	postConnection := PostConnection{
		edges: edges,
	}
	return postConnection
}

// Edges resolved
func (postConnection PostConnection) Edges() *[]*PostEdge {
	return postConnection.edges
}
