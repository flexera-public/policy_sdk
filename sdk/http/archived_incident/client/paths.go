// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the ArchivedIncident service.
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"fmt"
)

// ShowArchivedIncidentPath returns the URL path to the ArchivedIncident service show HTTP endpoint.
func ShowArchivedIncidentPath(projectID uint, incidentID string) string {
	return fmt.Sprintf("/api/governance/projects/%v/archived_incidents/%v", projectID, incidentID)
}

// IndexArchivedIncidentPath returns the URL path to the ArchivedIncident service index HTTP endpoint.
func IndexArchivedIncidentPath(projectID uint) string {
	return fmt.Sprintf("/api/governance/projects/%v/archived_incidents", projectID)
}

// IndexEscalationsArchivedIncidentPath returns the URL path to the ArchivedIncident service index_escalations HTTP endpoint.
func IndexEscalationsArchivedIncidentPath(projectID uint, incidentID string) string {
	return fmt.Sprintf("/api/governance/projects/%v/archived_incidents/%v/escalations", projectID, incidentID)
}

// IndexResolutionsArchivedIncidentPath returns the URL path to the ArchivedIncident service index_resolutions HTTP endpoint.
func IndexResolutionsArchivedIncidentPath(projectID uint, incidentID string) string {
	return fmt.Sprintf("/api/governance/projects/%v/archived_incidents/%v/resolutions", projectID, incidentID)
}