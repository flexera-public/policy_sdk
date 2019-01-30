// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// HTTP request path constructors for the PublishedTemplate service.
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"fmt"
)

// CreatePublishedTemplatePath returns the URL path to the PublishedTemplate service create HTTP endpoint.
func CreatePublishedTemplatePath(orgID uint) string {
	return fmt.Sprintf("/api/governance/orgs/%v/published_templates", orgID)
}

// UpdatePublishedTemplatePath returns the URL path to the PublishedTemplate service update HTTP endpoint.
func UpdatePublishedTemplatePath(orgID uint, templateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/published_templates/%v", orgID, templateID)
}

// HidePublishedTemplatePath returns the URL path to the PublishedTemplate service hide HTTP endpoint.
func HidePublishedTemplatePath(orgID uint, templateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/published_templates/%v/hide", orgID, templateID)
}

// UnhidePublishedTemplatePath returns the URL path to the PublishedTemplate service unhide HTTP endpoint.
func UnhidePublishedTemplatePath(orgID uint, templateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/published_templates/%v/unhide", orgID, templateID)
}

// DeletePublishedTemplatePath returns the URL path to the PublishedTemplate service delete HTTP endpoint.
func DeletePublishedTemplatePath(orgID uint, templateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/published_templates/%v", orgID, templateID)
}

// ShowPublishedTemplatePath returns the URL path to the PublishedTemplate service show HTTP endpoint.
func ShowPublishedTemplatePath(orgID uint, templateID string) string {
	return fmt.Sprintf("/api/governance/orgs/%v/published_templates/%v", orgID, templateID)
}

// IndexPublishedTemplatePath returns the URL path to the PublishedTemplate service index HTTP endpoint.
func IndexPublishedTemplatePath(orgID uint) string {
	return fmt.Sprintf("/api/governance/orgs/%v/published_templates", orgID)
}