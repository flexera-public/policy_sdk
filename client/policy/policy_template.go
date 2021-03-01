// Copyright (c) 2018 RightScale, Inc. - see LICENSE

package policy

import (
	"context"

	"github.com/pkg/errors"
	policytemplate "github.com/rightscale/policy_sdk/sdk/policy_template"
)

// UploadPolicyTemplate a policy template
func (c *client) UploadPolicyTemplate(ctx context.Context, filename, source string) (*policytemplate.PolicyTemplate, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &policytemplate.UploadPayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		Filename:   filename,
		Source:     source,
	}
	ptIF, err := c.pte.upload(ctx, p)
	if err != nil {
		return nil, err
	}
	pt, ok := ptIF.(*policytemplate.PolicyTemplate)
	if !ok {
		return nil, errors.New("error interpreting template")
	}
	return pt, nil
}

// UpdatePolicyTemplate a policy template
func (c *client) UpdatePolicyTemplate(ctx context.Context, id, filename, source string) (*policytemplate.PolicyTemplate, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &policytemplate.UpdatePayload{
		Token:      token,
		TemplateID: id,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		Filename:   filename,
		Source:     source,
	}
	ptIF, err := c.pte.update(ctx, p)
	if err != nil {
		return nil, err
	}
	pt, ok := ptIF.(*policytemplate.PolicyTemplate)
	if !ok {
		return nil, errors.New("error interpreting template")
	}
	return pt, nil
}

// CompilePolicyTemplate a policy template
func (c *client) CompilePolicyTemplate(ctx context.Context, filename, source string) error {
	token, err := c.getToken()
	if err != nil {
		return err
	}
	p := &policytemplate.CompilePayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		Filename:   filename,
		Source:     source,
	}
	_, err = c.pte.compile(ctx, p)
	return err
}

// DeletePolicyTemplate a policy template
func (c *client) DeletePolicyTemplate(ctx context.Context, templateID string) error {
	token, err := c.getToken()
	if err != nil {
		return err
	}
	p := &policytemplate.DeletePayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		TemplateID: templateID,
	}
	_, err = c.pte.delete(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

// ShowPolicyTemplate a policy template
func (c *client) ShowPolicyTemplate(ctx context.Context, templateID, view string) (*policytemplate.PolicyTemplate, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &policytemplate.ShowPayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		TemplateID: templateID,
	}

	if view != "" {
		p.View = &view
	}
	ptIF, err := c.pte.show(ctx, p)
	if err != nil {
		return nil, err
	}
	pt, ok := ptIF.(*policytemplate.PolicyTemplate)
	if !ok {
		return nil, errors.New("error interpreting template")
	}
	return pt, nil
}

// RetrieveData retrieves the  data sources defined in policy  template
func (c *client) RetrieveData(ctx context.Context, templateID string, names []string, options []*policytemplate.ConfigurationOptionCreateType, credentials map[string]string) ([]*policytemplate.Data, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}

	p := &policytemplate.RetrieveDataPayload{
		Token:       token,
		ProjectID:   c.projectID,
		Names:       names,
		TemplateID:  templateID,
		Options:     options,
		Credentials: credentials,
		APIVersion:  apiVersion,
	}

	resp, err := c.pte.retrieveData(ctx, p)
	if err != nil {
		return nil, err
	}
	data, ok := resp.([]*policytemplate.Data)
	if !ok {
		return nil, errors.New("error interpreting template")
	}
	return data, nil
}

func strPtr(s string) *string { return &s }
func boolPtr(b bool) *bool    { return &b }
