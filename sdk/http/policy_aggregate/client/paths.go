// Code generated by goa v3.1.3, DO NOT EDIT.
//
// HTTP request path constructors for the PolicyAggregate service.
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"fmt"
)

// CreatePolicyAggregatePath returns the URL path to the PolicyAggregate service create HTTP endpoint.
func CreatePolicyAggregatePath(orgID uint) string {
	return fmt.Sprintf("/api/governance/orgs/%v/policy_aggregates", orgID)
}

// UpdatePolicyAggregatePath returns the URL path to the PolicyAggregate service update HTTP endpoint.
func UpdatePolicyAggregatePath(orgID uint, policyAggregateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/policy_aggregates/%v", orgID, policyAggregateID)
}

// DeletePolicyAggregatePath returns the URL path to the PolicyAggregate service delete HTTP endpoint.
func DeletePolicyAggregatePath(orgID uint, policyAggregateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/policy_aggregates/%v", orgID, policyAggregateID)
}

// ShowPolicyAggregatePath returns the URL path to the PolicyAggregate service show HTTP endpoint.
func ShowPolicyAggregatePath(orgID uint, policyAggregateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/policy_aggregates/%v", orgID, policyAggregateID)
}

// ShowNonCatalogPolicyAggregatePath returns the URL path to the PolicyAggregate service show_non_catalog HTTP endpoint.
func ShowNonCatalogPolicyAggregatePath(orgID uint) string {
	return fmt.Sprintf("/api/governance/orgs/%v/policy_aggregates/non_catalog", orgID)
}

// IndexPolicyAggregatePath returns the URL path to the PolicyAggregate service index HTTP endpoint.
func IndexPolicyAggregatePath(orgID uint) string {
	return fmt.Sprintf("/api/governance/orgs/%v/policy_aggregates", orgID)
}
