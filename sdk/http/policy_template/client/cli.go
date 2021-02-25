// Code generated by goa v3.1.3, DO NOT EDIT.
//
// PolicyTemplate HTTP client CLI support package
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design -o .

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	policytemplate "github.com/rightscale/policy_sdk/sdk/policy_template"
	goa "goa.design/goa/v3/pkg"
)

// BuildCompilePayload builds the payload for the PolicyTemplate compile
// endpoint from CLI flags.
func BuildCompilePayload(policyTemplateCompileBody string, policyTemplateCompileProjectID string, policyTemplateCompileAPIVersion string, policyTemplateCompileToken string) (*policytemplate.CompilePayload, error) {
	var err error
	var body CompileRequestBody
	{
		err = json.Unmarshal([]byte(policyTemplateCompileBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"filename\": \"file.pt\",\n      \"source\": \"policy unattached_volumes do\\n\\t\\t\\t\\t\\t# ...\\n\\t\\t\\t\\t\\tend\"\n   }'")
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(policyTemplateCompileProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var apiVersion string
	{
		apiVersion = policyTemplateCompileAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if policyTemplateCompileToken != "" {
			token = &policyTemplateCompileToken
		}
	}
	v := &policytemplate.CompilePayload{
		Filename: body.Filename,
		Source:   body.Source,
	}
	v.ProjectID = projectID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildUploadPayload builds the payload for the PolicyTemplate upload endpoint
// from CLI flags.
func BuildUploadPayload(policyTemplateUploadBody string, policyTemplateUploadProjectID string, policyTemplateUploadAPIVersion string, policyTemplateUploadToken string) (*policytemplate.UploadPayload, error) {
	var err error
	var body UploadRequestBody
	{
		err = json.Unmarshal([]byte(policyTemplateUploadBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"filename\": \"file.pt\",\n      \"source\": \"policy unattached_volumes do\\n\\t\\t\\t\\t\\t# ...\\n\\t\\t\\t\\t\\tend\"\n   }'")
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(policyTemplateUploadProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var apiVersion string
	{
		apiVersion = policyTemplateUploadAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if policyTemplateUploadToken != "" {
			token = &policyTemplateUploadToken
		}
	}
	v := &policytemplate.UploadPayload{
		Filename: body.Filename,
		Source:   body.Source,
	}
	v.ProjectID = projectID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildUpdatePayload builds the payload for the PolicyTemplate update endpoint
// from CLI flags.
func BuildUpdatePayload(policyTemplateUpdateBody string, policyTemplateUpdateProjectID string, policyTemplateUpdateTemplateID string, policyTemplateUpdateAPIVersion string, policyTemplateUpdateToken string) (*policytemplate.UpdatePayload, error) {
	var err error
	var body UpdateRequestBody
	{
		err = json.Unmarshal([]byte(policyTemplateUpdateBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"filename\": \"tag_checker.pt\",\n      \"source\": \"policy unattached_volumes do\\n\\t\\t\\t\\t\\t# ...\\n\\t\\t\\t\\t\\tend\"\n   }'")
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(policyTemplateUpdateProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var templateID string
	{
		templateID = policyTemplateUpdateTemplateID
	}
	var apiVersion string
	{
		apiVersion = policyTemplateUpdateAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if policyTemplateUpdateToken != "" {
			token = &policyTemplateUpdateToken
		}
	}
	v := &policytemplate.UpdatePayload{
		Filename: body.Filename,
		Source:   body.Source,
	}
	v.ProjectID = projectID
	v.TemplateID = templateID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildRetrieveDataPayload builds the payload for the PolicyTemplate
// retrieve_data endpoint from CLI flags.
func BuildRetrieveDataPayload(policyTemplateRetrieveDataBody string, policyTemplateRetrieveDataProjectID string, policyTemplateRetrieveDataTemplateID string, policyTemplateRetrieveDataAPIVersion string, policyTemplateRetrieveDataToken string) (*policytemplate.RetrieveDataPayload, error) {
	var err error
	var body RetrieveDataRequestBody
	{
		err = json.Unmarshal([]byte(policyTemplateRetrieveDataBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"credentials\": {\n         \"Ea amet et.\": \"Illum possimus perferendis vitae suscipit.\"\n      },\n      \"names\": [\n         \"azure_resources\"\n      ],\n      \"options\": [\n         {\n            \"name\": \"cloud_vendor\",\n            \"value\": \"AWS\"\n         },\n         {\n            \"name\": \"email_list\",\n            \"value\": [\n               \"person1@domain.com\",\n               \"person2@domain.com\"\n            ]\n         }\n      ]\n   }'")
		}
	}
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(policyTemplateRetrieveDataProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var templateID string
	{
		templateID = policyTemplateRetrieveDataTemplateID
	}
	var apiVersion string
	{
		apiVersion = policyTemplateRetrieveDataAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if policyTemplateRetrieveDataToken != "" {
			token = &policyTemplateRetrieveDataToken
		}
	}
	v := &policytemplate.RetrieveDataPayload{}
	if body.Options != nil {
		v.Options = make([]*policytemplate.ConfigurationOptionCreateType, len(body.Options))
		for i, val := range body.Options {
			v.Options[i] = marshalConfigurationOptionCreateTypeRequestBodyToPolicytemplateConfigurationOptionCreateType(val)
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
	if body.Names != nil {
		v.Names = make([]string, len(body.Names))
		for i, val := range body.Names {
			v.Names[i] = val
		}
	}
	v.ProjectID = projectID
	v.TemplateID = templateID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildShowPayload builds the payload for the PolicyTemplate show endpoint
// from CLI flags.
func BuildShowPayload(policyTemplateShowProjectID string, policyTemplateShowTemplateID string, policyTemplateShowView string, policyTemplateShowAPIVersion string, policyTemplateShowToken string) (*policytemplate.ShowPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(policyTemplateShowProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var templateID string
	{
		templateID = policyTemplateShowTemplateID
	}
	var view *string
	{
		if policyTemplateShowView != "" {
			view = &policyTemplateShowView
			if view != nil {
				if !(*view == "default" || *view == "extended" || *view == "source") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "extended", "source"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var apiVersion string
	{
		apiVersion = policyTemplateShowAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if policyTemplateShowToken != "" {
			token = &policyTemplateShowToken
		}
	}
	v := &policytemplate.ShowPayload{}
	v.ProjectID = projectID
	v.TemplateID = templateID
	v.View = view
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildIndexPayload builds the payload for the PolicyTemplate index endpoint
// from CLI flags.
func BuildIndexPayload(policyTemplateIndexProjectID string, policyTemplateIndexView string, policyTemplateIndexAPIVersion string, policyTemplateIndexEtag string, policyTemplateIndexToken string) (*policytemplate.IndexPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(policyTemplateIndexProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var view *string
	{
		if policyTemplateIndexView != "" {
			view = &policyTemplateIndexView
			if view != nil {
				if !(*view == "default" || *view == "extended") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "extended"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var apiVersion string
	{
		apiVersion = policyTemplateIndexAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var etag *string
	{
		if policyTemplateIndexEtag != "" {
			etag = &policyTemplateIndexEtag
		}
	}
	var token *string
	{
		if policyTemplateIndexToken != "" {
			token = &policyTemplateIndexToken
		}
	}
	v := &policytemplate.IndexPayload{}
	v.ProjectID = projectID
	v.View = view
	v.APIVersion = apiVersion
	v.Etag = etag
	v.Token = token

	return v, nil
}

// BuildDeletePayload builds the payload for the PolicyTemplate delete endpoint
// from CLI flags.
func BuildDeletePayload(policyTemplateDeleteProjectID string, policyTemplateDeleteTemplateID string, policyTemplateDeleteAPIVersion string, policyTemplateDeleteToken string) (*policytemplate.DeletePayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(policyTemplateDeleteProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var templateID string
	{
		templateID = policyTemplateDeleteTemplateID
	}
	var apiVersion string
	{
		apiVersion = policyTemplateDeleteAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if policyTemplateDeleteToken != "" {
			token = &policyTemplateDeleteToken
		}
	}
	v := &policytemplate.DeletePayload{}
	v.ProjectID = projectID
	v.TemplateID = templateID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}
