package query

import (
	"strings"

	"github.com/RoyalIcing/collected-systems/types"
)

// FilterPeople uses args to filter a list of all people down
func FilterPeople(allPeople []*types.Person, args PeopleArgs) (*[]*types.Person, error) {
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
