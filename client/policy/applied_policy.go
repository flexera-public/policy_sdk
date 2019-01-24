// Copyright (c) 2018 RightScale, Inc. - see LICENSE

package policy

import (
	"context"

	"github.com/pkg/errors"
	"github.com/rightscale/right_pt/sdk/applied_policy"
)

// CreateAppliedPolicy an applied policy
func (c *client) CreateAppliedPolicy(ctx context.Context, p *appliedpolicy.CreatePayload) (*appliedpolicy.AppliedPolicy, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p.Token = token
	p.APIVersion = apiVersion
	p.ProjectID = c.projectID

	policyIF, err := c.ape.create(ctx, p)
	if err != nil {
		return nil, err
	}
	policy, ok := policyIF.(*appliedpolicy.AppliedPolicy)
	if !ok {
		return nil, errors.New("error interpreting applied policy")
	}
	return policy, nil
}

// DeleteAppliedPolicy deletes an applied policy
func (c *client) DeleteAppliedPolicy(ctx context.Context, id string) error {
	token, err := c.getToken()
	if err != nil {
		return err
	}
	p := &appliedpolicy.DeletePayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		PolicyID:   id,
	}
	_, err = c.ape.delete(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

// ShowAppliedPolicy shows an applied policy
func (c *client) ShowAppliedPolicy(ctx context.Context, id, view string) (*appliedpolicy.AppliedPolicy, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &appliedpolicy.ShowPayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		PolicyID:   id,
	}
	if view != "" {
		p.View = &view
	}
	policyIF, err := c.ape.show(ctx, p)
	if err != nil {
		return nil, err
	}
	policy, ok := policyIF.(*appliedpolicy.AppliedPolicy)
	if !ok {
		return nil, errors.New("error interpreting applied policy")

	}
	return policy, nil
}

// IndexAppliedPolicies returns a list of applied policies with optional filtering
func (c *client) IndexAppliedPolicies(ctx context.Context, names []string, view, etag string) (*appliedpolicy.AppliedPolicyList, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &appliedpolicy.IndexPayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		Name:       names,
	}
	if view != "" {
		p.View = &view
	}
	if etag != "" {
		p.Etag = &etag
	}
	apListIF, err := c.ape.index(ctx, p)
	if err != nil {
		return nil, err
	}
	apList, ok := apListIF.(*appliedpolicy.AppliedPolicyList)
	if !ok {
		return nil, errors.New("error interpreting applied policy list")

	}
	return apList, nil
}
