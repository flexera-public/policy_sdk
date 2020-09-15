// Code generated by goa v3.1.3, DO NOT EDIT.
//
// AppliedPolicy HTTP client CLI support package
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	appliedpolicy "github.com/rightscale/policy_sdk/sdk/applied_policy"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreatePayload builds the payload for the AppliedPolicy create endpoint
// from CLI flags.
func BuildCreatePayload(appliedPolicyCreateBody string, appliedPolicyCreateProjectID string, appliedPolicyCreateAPIVersion string, appliedPolicyCreateToken string) (*appliedpolicy.CreatePayload, error) {
	var err error
	var body CreateRequestBody
	{
		err = json.Unmarshal([]byte(appliedPolicyCreateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"credentials\": {\n         \"cred_aws\": \"aws_default\"\n      },\n      \"description\": \"Delete unattached volumes after 24 hours in US-East.\",\n      \"dry_run\": false,\n      \"frequency\": \"daily\",\n      \"name\": \"us_east_unattached_volumes\",\n      \"options\": [\n         {\n            \"name\": \"cloud_vendor\",\n            \"value\": \"AWS\"\n         },\n         {\n            \"name\": \"email_list\",\n            \"value\": [\n               \"person1@domain.com\",\n               \"person2@domain.com\"\n            ]\n         }\n      ],\n      \"severity\": \"low\",\n      \"skip_approvals\": false,\n      \"template_href\": \"/api/governance/projects/60073/policy_templates/5b06ead5e0dacc007058c784\"\n   }'")
		}
		err = goa.MergeErrors(err, goa.ValidatePattern("body.template_href", body.TemplateHref, "^/api/governance/(projects/[0-9]+/policy|orgs/[0-9]+/published)_templates/[0-9a-f]+$"))
		if !(body.Frequency == "15 minutes" || body.Frequency == "hourly" || body.Frequency == "daily" || body.Frequency == "weekly" || body.Frequency == "monthly") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.frequency", body.Frequency, []interface{}{"15 minutes", "hourly", "daily", "weekly", "monthly"}))
		}
		if body.Severity != nil {
			if !(*body.Severity == "low" || *body.Severity == "medium" || *body.Severity == "high" || *body.Severity == "critical") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.severity", *body.Severity, []interface{}{"low", "medium", "high", "critical"}))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyCreateProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyCreateAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if appliedPolicyCreateToken != "" {
			token = &appliedPolicyCreateToken
		}
	}
	v := &appliedpolicy.CreatePayload{
		Name:          body.Name,
		Description:   body.Description,
		TemplateHref:  body.TemplateHref,
		Frequency:     body.Frequency,
		DryRun:        body.DryRun,
		SkipApprovals: body.SkipApprovals,
		Severity:      body.Severity,
	}
	if body.Options != nil {
		v.Options = make([]*appliedpolicy.ConfigurationOptionCreateType, len(body.Options))
		for i, val := range body.Options {
			v.Options[i] = marshalConfigurationOptionCreateTypeRequestBodyToAppliedpolicyConfigurationOptionCreateType(val)
		}
	}
	if body.Credentials != nil {
		v.Credentials = make(map[string]string, len(body.Credentials))
		for key, val := range body.Credentials {
			tk := key
			tv := val
			v.Credentials[tk] = tv
		}
	}
	v.ProjectID = projectID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildDeletePayload builds the payload for the AppliedPolicy delete endpoint
// from CLI flags.
func BuildDeletePayload(appliedPolicyDeleteProjectID string, appliedPolicyDeletePolicyID string, appliedPolicyDeleteAPIVersion string, appliedPolicyDeleteToken string) (*appliedpolicy.DeletePayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyDeleteProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var policyID string
	{
		policyID = appliedPolicyDeletePolicyID
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyDeleteAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if appliedPolicyDeleteToken != "" {
			token = &appliedPolicyDeleteToken
		}
	}
	v := &appliedpolicy.DeletePayload{}
	v.ProjectID = projectID
	v.PolicyID = policyID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildShowPayload builds the payload for the AppliedPolicy show endpoint from
// CLI flags.
func BuildShowPayload(appliedPolicyShowProjectID string, appliedPolicyShowPolicyID string, appliedPolicyShowView string, appliedPolicyShowAPIVersion string, appliedPolicyShowToken string) (*appliedpolicy.ShowPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyShowProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var policyID string
	{
		policyID = appliedPolicyShowPolicyID
	}
	var view *string
	{
		if appliedPolicyShowView != "" {
			view = &appliedPolicyShowView
			if view != nil {
				if !(*view == "default" || *view == "source") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "source"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyShowAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if appliedPolicyShowToken != "" {
			token = &appliedPolicyShowToken
		}
	}
	v := &appliedpolicy.ShowPayload{}
	v.ProjectID = projectID
	v.PolicyID = policyID
	v.View = view
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildShowStatusPayload builds the payload for the AppliedPolicy show_status
// endpoint from CLI flags.
func BuildShowStatusPayload(appliedPolicyShowStatusProjectID string, appliedPolicyShowStatusPolicyID string, appliedPolicyShowStatusAPIVersion string, appliedPolicyShowStatusToken string) (*appliedpolicy.ShowStatusPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyShowStatusProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var policyID string
	{
		policyID = appliedPolicyShowStatusPolicyID
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyShowStatusAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if appliedPolicyShowStatusToken != "" {
			token = &appliedPolicyShowStatusToken
		}
	}
	v := &appliedpolicy.ShowStatusPayload{}
	v.ProjectID = projectID
	v.PolicyID = policyID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildShowLogPayload builds the payload for the AppliedPolicy show_log
// endpoint from CLI flags.
func BuildShowLogPayload(appliedPolicyShowLogProjectID string, appliedPolicyShowLogPolicyID string, appliedPolicyShowLogAPIVersion string, appliedPolicyShowLogEtag string, appliedPolicyShowLogToken string) (*appliedpolicy.ShowLogPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyShowLogProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var policyID string
	{
		policyID = appliedPolicyShowLogPolicyID
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyShowLogAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var etag *string
	{
		if appliedPolicyShowLogEtag != "" {
			etag = &appliedPolicyShowLogEtag
		}
	}
	var token *string
	{
		if appliedPolicyShowLogToken != "" {
			token = &appliedPolicyShowLogToken
		}
	}
	v := &appliedpolicy.ShowLogPayload{}
	v.ProjectID = projectID
	v.PolicyID = policyID
	v.APIVersion = apiVersion
	v.Etag = etag
	v.Token = token

	return v, nil
}

// BuildIndexPayload builds the payload for the AppliedPolicy index endpoint
// from CLI flags.
func BuildIndexPayload(appliedPolicyIndexProjectID string, appliedPolicyIndexView string, appliedPolicyIndexName string, appliedPolicyIndexAPIVersion string, appliedPolicyIndexEtag string, appliedPolicyIndexToken string) (*appliedpolicy.IndexPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyIndexProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var view *string
	{
		if appliedPolicyIndexView != "" {
			view = &appliedPolicyIndexView
			if view != nil {
				if !(*view == "default") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var name []string
	{
		if appliedPolicyIndexName != "" {
			err = json.Unmarshal([]byte(appliedPolicyIndexName), &name)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for name, example of valid JSON:\n%s", "'[\n      \"Tag Checker Policy\"\n   ]'")
			}
		}
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyIndexAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var etag *string
	{
		if appliedPolicyIndexEtag != "" {
			etag = &appliedPolicyIndexEtag
		}
	}
	var token *string
	{
		if appliedPolicyIndexToken != "" {
			token = &appliedPolicyIndexToken
		}
	}
	v := &appliedpolicy.IndexPayload{}
	v.ProjectID = projectID
	v.View = view
	v.Name = name
	v.APIVersion = apiVersion
	v.Etag = etag
	v.Token = token

	return v, nil
}

// BuildEvaluatePayload builds the payload for the AppliedPolicy evaluate
// endpoint from CLI flags.
func BuildEvaluatePayload(appliedPolicyEvaluateProjectID string, appliedPolicyEvaluatePolicyID string, appliedPolicyEvaluateAPIVersion string, appliedPolicyEvaluateToken string) (*appliedpolicy.EvaluatePayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyEvaluateProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var policyID string
	{
		policyID = appliedPolicyEvaluatePolicyID
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyEvaluateAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if appliedPolicyEvaluateToken != "" {
			token = &appliedPolicyEvaluateToken
		}
	}
	v := &appliedpolicy.EvaluatePayload{}
	v.ProjectID = projectID
	v.PolicyID = policyID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildUpdatePayload builds the payload for the AppliedPolicy update endpoint
// from CLI flags.
func BuildUpdatePayload(appliedPolicyUpdateBody string, appliedPolicyUpdateProjectID string, appliedPolicyUpdatePolicyID string, appliedPolicyUpdateAPIVersion string, appliedPolicyUpdateToken string) (*appliedpolicy.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(appliedPolicyUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"description\": \"Delete unattached volumes after 24 hours in US-East.\",\n      \"dry_run\": false,\n      \"frequency\": \"daily\",\n      \"name\": \"us_east_unattached_volumes\",\n      \"options\": [\n         {\n            \"name\": \"cloud_vendor\",\n            \"value\": \"AWS\"\n         },\n         {\n            \"name\": \"email_list\",\n            \"value\": [\n               \"person1@domain.com\",\n               \"person2@domain.com\"\n            ]\n         }\n      ],\n      \"severity\": \"low\",\n      \"skip_approvals\": false\n   }'")
		}
		if body.Frequency != nil {
			if !(*body.Frequency == "15 minutes" || *body.Frequency == "hourly" || *body.Frequency == "daily" || *body.Frequency == "weekly" || *body.Frequency == "monthly") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.frequency", *body.Frequency, []interface{}{"15 minutes", "hourly", "daily", "weekly", "monthly"}))
			}
		}
		if body.Severity != nil {
			if !(*body.Severity == "low" || *body.Severity == "medium" || *body.Severity == "high" || *body.Severity == "critical") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.severity", *body.Severity, []interface{}{"low", "medium", "high", "critical"}))
			}
		}
		if err != nil {
			return nil, err
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(appliedPolicyUpdateProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var policyID string
	{
		policyID = appliedPolicyUpdatePolicyID
	}
	var apiVersion string
	{
		apiVersion = appliedPolicyUpdateAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if appliedPolicyUpdateToken != "" {
			token = &appliedPolicyUpdateToken
		}
	}
	v := &appliedpolicy.UpdatePayload{
		Name:          body.Name,
		Description:   body.Description,
		Frequency:     body.Frequency,
		DryRun:        body.DryRun,
		SkipApprovals: body.SkipApprovals,
		Severity:      body.Severity,
	}
	if body.Options != nil {
		v.Options = make([]*appliedpolicy.ConfigurationOptionCreateType, len(body.Options))
		for i, val := range body.Options {
			v.Options[i] = marshalConfigurationOptionCreateTypeRequestBodyToAppliedpolicyConfigurationOptionCreateType(val)
		}
	}
	v.ProjectID = projectID
	v.PolicyID = policyID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}
