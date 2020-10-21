// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ArchivedIncident HTTP client CLI support package
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	archivedincident "github.com/rightscale/policy_sdk/sdk/archived_incident"
	goa "goa.design/goa/v3/pkg"
)

// BuildShowPayload builds the payload for the ArchivedIncident show endpoint
// from CLI flags.
func BuildShowPayload(archivedIncidentShowProjectID string, archivedIncidentShowIncidentID string, archivedIncidentShowView string, archivedIncidentShowAPIVersion string, archivedIncidentShowEtag string, archivedIncidentShowToken string) (*archivedincident.ShowPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(archivedIncidentShowProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var incidentID string
	{
		incidentID = archivedIncidentShowIncidentID
	}
	var view *string
	{
		if archivedIncidentShowView != "" {
			view = &archivedIncidentShowView
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
		apiVersion = archivedIncidentShowAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var etag *string
	{
		if archivedIncidentShowEtag != "" {
			etag = &archivedIncidentShowEtag
		}
	}
	var token *string
	{
		if archivedIncidentShowToken != "" {
			token = &archivedIncidentShowToken
		}
	}
	v := &archivedincident.ShowPayload{}
	v.ProjectID = projectID
	v.IncidentID = incidentID
	v.View = view
	v.APIVersion = apiVersion
	v.Etag = etag
	v.Token = token

	return v, nil
}

// BuildIndexPayload builds the payload for the ArchivedIncident index endpoint
// from CLI flags.
func BuildIndexPayload(archivedIncidentIndexProjectID string, archivedIncidentIndexView string, archivedIncidentIndexState string, archivedIncidentIndexAppliedPolicyID string, archivedIncidentIndexAPIVersion string, archivedIncidentIndexEtag string, archivedIncidentIndexToken string) (*archivedincident.IndexPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(archivedIncidentIndexProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var view *string
	{
		if archivedIncidentIndexView != "" {
			view = &archivedIncidentIndexView
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
	var state []string
	{
		if archivedIncidentIndexState != "" {
			err = json.Unmarshal([]byte(archivedIncidentIndexState), &state)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for state, example of valid JSON:\n%s", "'[\n      \"terminated\"\n   ]'")
			}
			for _, e := range state {
				if !(e == "resolved" || e == "terminated") {
					err = goa.MergeErrors(err, goa.InvalidEnumValueError("state[*]", e, []interface{}{"resolved", "terminated"}))
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	var appliedPolicyID *string
	{
		if archivedIncidentIndexAppliedPolicyID != "" {
			appliedPolicyID = &archivedIncidentIndexAppliedPolicyID
		}
	}
	var apiVersion string
	{
		apiVersion = archivedIncidentIndexAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var etag *string
	{
		if archivedIncidentIndexEtag != "" {
			etag = &archivedIncidentIndexEtag
		}
	}
	var token *string
	{
		if archivedIncidentIndexToken != "" {
			token = &archivedIncidentIndexToken
		}
	}
	v := &archivedincident.IndexPayload{}
	v.ProjectID = projectID
	v.View = view
	v.State = state
	v.AppliedPolicyID = appliedPolicyID
	v.APIVersion = apiVersion
	v.Etag = etag
	v.Token = token

	return v, nil
}

// BuildIndexEscalationsPayload builds the payload for the ArchivedIncident
// index_escalations endpoint from CLI flags.
func BuildIndexEscalationsPayload(archivedIncidentIndexEscalationsProjectID string, archivedIncidentIndexEscalationsIncidentID string, archivedIncidentIndexEscalationsAPIVersion string, archivedIncidentIndexEscalationsToken string) (*archivedincident.IndexEscalationsPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(archivedIncidentIndexEscalationsProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var incidentID string
	{
		incidentID = archivedIncidentIndexEscalationsIncidentID
	}
	var apiVersion string
	{
		apiVersion = archivedIncidentIndexEscalationsAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if archivedIncidentIndexEscalationsToken != "" {
			token = &archivedIncidentIndexEscalationsToken
		}
	}
	v := &archivedincident.IndexEscalationsPayload{}
	v.ProjectID = projectID
	v.IncidentID = incidentID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}

// BuildIndexResolutionsPayload builds the payload for the ArchivedIncident
// index_resolutions endpoint from CLI flags.
func BuildIndexResolutionsPayload(archivedIncidentIndexResolutionsProjectID string, archivedIncidentIndexResolutionsIncidentID string, archivedIncidentIndexResolutionsAPIVersion string, archivedIncidentIndexResolutionsToken string) (*archivedincident.IndexResolutionsPayload, error) {
	var err error
	var projectID uint
	{
		var v uint64
		v, err = strconv.ParseUint(archivedIncidentIndexResolutionsProjectID, 10, 64)
		projectID = uint(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for projectID, must be UINT")
		}
	}
	var incidentID string
	{
		incidentID = archivedIncidentIndexResolutionsIncidentID
	}
	var apiVersion string
	{
		apiVersion = archivedIncidentIndexResolutionsAPIVersion
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		if err != nil {
			return nil, err
		}
	}
	var token *string
	{
		if archivedIncidentIndexResolutionsToken != "" {
			token = &archivedIncidentIndexResolutionsToken
		}
	}
	v := &archivedincident.IndexResolutionsPayload{}
	v.ProjectID = projectID
	v.IncidentID = incidentID
	v.APIVersion = apiVersion
	v.Token = token

	return v, nil
}
