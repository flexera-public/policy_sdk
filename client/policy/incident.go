// Copyright (c) 2018 RightScale, Inc. - see LICENSE

package policy

import (
	"context"

	"github.com/pkg/errors"
	"github.com/flexera-public/policy_sdk/sdk/incident"
)

// ShowIncident an incident
func (c *client) ShowIncident(ctx context.Context, id string, view string) (*incident.Incident, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &incident.ShowPayload{
		Token:      token,
		ProjectID:  c.projectID,
		IncidentID: id,
		APIVersion: apiVersion,
	}
	if view != "" {
		p.View = &view
	}
	iIF, err := c.ie.show(ctx, p)
	if err != nil {
		return nil, err
	}
	i, ok := iIF.(*incident.Incident)
	if !ok {
		return nil, errors.New("error interpreting incident")
	}
	return i, nil

}

// IndexIncidents incidents
func (c *client) IndexIncidents(ctx context.Context, appliedPolicyID string, state []string, view string, eTag string) (*incident.IncidentList, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &incident.IndexPayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
	}
	if view != "" {
		p.View = &view
	}
	if eTag != "" {
		p.Etag = &eTag
	}
	if appliedPolicyID != "" {
		p.AppliedPolicyID = &appliedPolicyID
	}
	if len(state) > 0 {
		p.State = state
	}
	ilIF, err := c.ie.index(ctx, p)
	if err != nil {
		return nil, err
	}
	il, ok := ilIF.(*incident.IncidentList)
	if !ok {
		return nil, errors.New("error interpreting incident list")
	}
	return il, nil
}

// IndexEscalations for an incident
func (c *client) IndexEscalations(ctx context.Context, id string) (*incident.Escalations, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &incident.IndexEscalationsPayload{
		Token:      token,
		ProjectID:  c.projectID,
		IncidentID: id,
		APIVersion: apiVersion,
	}
	resIF, err := c.ie.escalations(ctx, p)
	if err != nil {
		return nil, err
	}
	res, ok := resIF.(*incident.Escalations)
	if !ok {
		return nil, errors.New("error interpreting escalations")
	}
	return res, nil
}
