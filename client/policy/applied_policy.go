// Copyright (c) 2018 RightScale, Inc. - see LICENSE

package policy

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	appliedpolicy "github.com/flexera-public/policy_sdk/sdk/applied_policy"
	apclient "github.com/flexera-public/policy_sdk/sdk/http/applied_policy/client"
	"github.com/pkg/errors"
	goahttp "goa.design/goa/v3/http"
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

// ShowAppliedPolicyLog shows an applied policy
func (c *client) ShowAppliedPolicyLog(ctx context.Context, id, etag string) (*appliedpolicy.AppliedPolicyLog, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &appliedpolicy.ShowLogPayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		PolicyID:   id,
	}
	if etag != "" {
		p.Etag = &etag
	}
	logIF, err := c.ape.showLog(ctx, p)
	if err != nil {
		return nil, err
	}
	log, ok := logIF.(*appliedpolicy.AppliedPolicyLog)
	if !ok {
		return nil, errors.New("error interpreting applied policy log")

	}
	return log, nil
}

func showLogResponseDecoder(resp *http.Response) goahttp.Decoder {
	return &showLogDecoder{resp.Body}
}

type showLogDecoder struct {
	body io.Reader
}

func (d *showLogDecoder) Decode(v interface{}) error {
	bodyBytes, err := ioutil.ReadAll(d.body)
	if err != nil {
		return err
	}

	switch t := v.(type) {
	case *string:
		*t = string(bodyBytes)
	case *apclient.ShowLogNotFoundResponseBody:
		if len(bodyBytes) > 0 {
			json.Unmarshal(bodyBytes, t)
		} else {
			t.Name = strPtr("not_found")
			t.ID = strPtr("missing")
			t.Message = strPtr("log not found")
			t.Temporary = boolPtr(false)
			t.Timeout = boolPtr(false)
			t.Fault = boolPtr(false)
		}
	default:
		json.Unmarshal(bodyBytes, t)
	}
	return nil
}

// ShowAppliedPolicyLog shows an applied policy
// func (c *client) ShowAppliedPolicyLog(ctx context.Context, id, etag string) ([]byte, string, error) {
// 	token, err := c.getToken()
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	p := &appliedpolicy.ShowLogPayload{
// 		Token:      token,
// 		ProjectID:  c.projectID,
// 		APIVersion: apiVersion,
// 		PolicyID:   id,
// 	}
// 	if etag != "" {
// 		p.Etag = &etag
// 	}
// 	u := &url.URL{
// 		Scheme: "https",
// 		Host:   c.host,
// 		Path:   apclient.ShowLogAppliedPolicyPath(p.ProjectID, p.PolicyID)}
// 	req, err := http.NewRequest("GET", u.String(), nil)
// 	if err != nil {
// 		return nil, "", goahttp.ErrInvalidURL("AppliedPolicy", "log", u.String(), err)
// 	}
// 	if ctx != nil {
// 		req = req.WithContext(ctx)
// 	}

// 	var encodeRequest = apclient.EncodeShowLogRequest(goahttp.RequestEncoder)
// 	if err = encodeRequest(req, p); err != nil {
// 		return nil, "", err
// 	}

// 	resp, err := c.showLogDoer.Do(req)
// 	if err != nil {
// 		return nil, "", goahttp.ErrRequestError("AppliedPolicy", "log", err)
// 	}
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	etag = resp.Header.Get("ETag")
// 	return body, etag, nil
// }

// ShowAppliedPolicyStatus shows an applied policy
func (c *client) ShowAppliedPolicyStatus(ctx context.Context, id string) (*appliedpolicy.AppliedPolicyStatus, error) {
	token, err := c.getToken()
	if err != nil {
		return nil, err
	}
	p := &appliedpolicy.ShowStatusPayload{
		Token:      token,
		ProjectID:  c.projectID,
		APIVersion: apiVersion,
		PolicyID:   id,
	}
	statusIF, err := c.ape.showStatus(ctx, p)
	if err != nil {
		return nil, err
	}
	status, ok := statusIF.(*appliedpolicy.AppliedPolicyStatus)
	if !ok {
		return nil, errors.New("error interpreting applied policy")

	}
	return status, nil
}
