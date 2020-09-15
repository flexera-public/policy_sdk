// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the IncidentAggregate service.
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"fmt"
)

// ShowIncidentAggregatePath returns the URL path to the IncidentAggregate service show HTTP endpoint.
func ShowIncidentAggregatePath(orgID uint, incidentAggregateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/incident_aggregates/%v", orgID, incidentAggregateID)
}

// ShowNonCatalogIncidentAggregatePath returns the URL path to the IncidentAggregate service show_non_catalog HTTP endpoint.
func ShowNonCatalogIncidentAggregatePath(orgID uint) string {
	return fmt.Sprintf("/api/governance/orgs/%v/incident_aggregates/non_catalog", orgID)
}

// IndexIncidentAggregatePath returns the URL path to the IncidentAggregate service index HTTP endpoint.
func IndexIncidentAggregatePath(orgID uint) string {
	return fmt.Sprintf("/api/governance/orgs/%v/incident_aggregates", orgID)
}
