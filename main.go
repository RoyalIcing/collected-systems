package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/graph-gophers/graphql-go/relay"

	"github.com/RoyalIcing/collected-systems/query"
	"github.com/RoyalIcing/collected-systems/sources"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3838"
	}

	localFileSystemSource := sources.NewLocalFileSource(packr.NewBox("./samples/RoyalIcing"))
	schema := query.MakeSchema(&localFileSystemSource)

	http.Handle("/", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
