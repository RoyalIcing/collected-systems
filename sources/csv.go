package sources

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"github.com/RoyalIcing/collected-systems/types"
)

func parseRoles(input string) []string {
	elements := strings.Split(input, ",")
	adjusted := elements[:0]
	for _, s := range elements {
		adjusted = append(adjusted, strings.ToUpper(s))
	}
	return adjusted
}

var profileDomains = map[string]bool{
	"github.com":  true,
	"medium.com":  true,
	"twitter.com": true,
}

// ReadPeopleCSVFrom loads all the people stored in a CSV source
func ReadPeopleCSVFrom(reader io.Reader) (*[]*types.Person, error) {
	people := []*types.Person{}
	csvReader := csv.NewReader(reader)

	headers, err := csvReader.Read()
	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}

			firstName := ""
			lastName := ""
			var roles []string
			profileUsernames := make(map[string]string)

			for i, header := range headers {
				if header == "firstName" {
					firstName = record[i]
				} else if header == "lastName" {
					lastName = record[i]
				} else if header == "roles" {
					roles = parseRoles(record[i])
				} else if profileDomains[header] {
					s := record[i]
					if s != "" {
						profileUsernames[header] = s
					}
				}
			}

			person := types.NewPerson(firstName, lastName, roles, profileUsernames)
			people = append(people, person)
		}

		break
	}
	return &people, nil
}

// ReadPeopleCSVFile loads all the people stored in a CSV file
func ReadPeopleCSVFile(path string) (*[]*types.Person, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ReadPeopleCSVFrom(file)
}

// ReadServicesCSVFrom loads all the services stored in a CSV source
func ReadServicesCSVFrom(reader io.Reader) (*[]*types.Service, error) {
	services := []*types.Service{}
	csvReader := csv.NewReader(reader)

	headers, err := csvReader.Read()
	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}

			domain := ""

			for i, header := range headers {
				if header == "domain" {
					domain = record[i]
				}
			}

			if domain != "" {
				service := types.NewService(domain)
				services = append(services, service)
			}
		}

		break
	}
	return &services, nil
}

// ReadServicesCSVFile loads all the services stored in a CSV file
func ReadServicesCSVFile(path string) (*[]*types.Service, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ReadServicesCSVFrom(file)
}

// ReadPostsCSVFrom loads all the posts stored in a CSV source
func ReadPostsCSVFrom(reader io.Reader) (*[]*types.PostEdge, error) {
	postEdges := []*types.PostEdge{}
	csvReader := csv.NewReader(reader)

	headers, err := csvReader.Read()
	for {
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		for {
			record, err := csvReader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}

			id := ""
			markdownSource := ""

			for i, header := range headers {
				if header == "id" {
					id = record[i]
				} else if header == "markdownSource" {
					markdownSource = record[i]
				}
			}

			if id != "" && markdownSource != "" {
				markdownDocument := types.NewMarkdownDocument(markdownSource)

				post := types.NewPost(id, markdownDocument)

				postEdge := types.NewPostEdge(post, id)
				postEdges = append(postEdges, postEdge)
			}
		}

		break
	}
	return &postEdges, nil
}

// ReadPostsCSVFile loads all the posts stored in a CSV file
func ReadPostsCSVFile(path string) (*[]*types.PostEdge, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return ReadPostsCSVFrom(file)
}
