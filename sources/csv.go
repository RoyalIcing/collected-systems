package sources

import (
	"encoding/csv"
	"io"
	"os"
	"strings"

	"../types"
)

func parseRoles(input string) []string {
	elements := strings.Split(input, ",")
	adjusted := elements[:0]
	for _, s := range elements {
		adjusted = append(adjusted, strings.ToUpper(s))
	}
	return adjusted
}

func readPeopleCSVFrom(reader io.Reader) (*[]*types.Person, error) {
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
			var gitHubUsername *string

			for i, header := range headers {
				if header == "firstName" {
					firstName = record[i]
				} else if header == "lastName" {
					lastName = record[i]
				} else if header == "roles" {
					roles = parseRoles(record[i])
				} else if header == "gitHubUsername" {
					s := record[i]
					gitHubUsername = &s
				}
			}

			person := types.NewPerson(firstName, lastName, roles, gitHubUsername)
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

	return readPeopleCSVFrom(file)
}
