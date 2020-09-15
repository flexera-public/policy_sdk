// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the ActionStatus service.
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package server

import (
	"fmt"
)

// IndexActionStatusPath returns the URL path to the ActionStatus service index HTTP endpoint.
func IndexActionStatusPath(projectID uint) string {
	return fmt.Sprintf("/api/governance/projects/%v/action_status", projectID)
}

// ShowActionStatusPath returns the URL path to the ActionStatus service show HTTP endpoint.
func ShowActionStatusPath(projectID uint, id string) string {
	return fmt.Sprintf("/api/governance/projects/%v/action_status/%v", projectID, id)
}
