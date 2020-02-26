// Code generated by goa v3.0.6, DO NOT EDIT.
//
// ArchivedIncident HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package server

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	archivedincidentviews "github.com/rightscale/governance/front_service/gen/archived_incident/views"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeShowResponse returns an encoder for responses returned by the
// ArchivedIncident show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*archivedincidentviews.ArchivedIncident)
		w.Header().Set("goa-view", res.View)
		if res.Projected.NotModified != nil && *res.Projected.NotModified == "true" {
			w.Header().Set("Etag", *res.Projected.Etag)
			w.WriteHeader(http.StatusNotModified)
			return nil
		}
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewShowOKResponseBody(res.Projected)
		case "extended":
			body = NewShowOKResponseBodyExtended(res.Projected)
		case "source":
			body = NewShowOKResponseBodySource(res.Projected)
		}
		if res.Projected.Etag != nil {
			w.Header().Set("Etag", *res.Projected.Etag)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the
// ArchivedIncident show endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			incidentID string
			view       *string
			apiVersion string
			etag       *string
			token      *string
			err        error

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		incidentID = params["incident_id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "extended" || *view == "source") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "extended", "source"}))
			}
		}
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		etagRaw := r.Header.Get("If-None-Match")
		if etagRaw != "" {
			etag = &etagRaw
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(projectID, incidentID, view, apiVersion, etag, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show
// ArchivedIncident endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewShowInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeIndexResponse returns an encoder for responses returned by the
// ArchivedIncident index endpoint.
func EncodeIndexResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*archivedincidentviews.ArchivedIncidentList)
		w.Header().Set("goa-view", res.View)
		if res.Projected.NotModified != nil && *res.Projected.NotModified == "true" {
			w.Header().Set("Etag", *res.Projected.Etag)
			w.WriteHeader(http.StatusNotModified)
			return nil
		}
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewIndexOKResponseBody(res.Projected)
		case "extended":
			body = NewIndexOKResponseBodyExtended(res.Projected)
		}
		if res.Projected.Etag != nil {
			w.Header().Set("Etag", *res.Projected.Etag)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeIndexRequest returns a decoder for requests sent to the
// ArchivedIncident index endpoint.
func DecodeIndexRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID       uint
			view            *string
			state           []string
			appliedPolicyID *string
			apiVersion      string
			etag            *string
			token           *string
			err             error

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "extended") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "extended"}))
			}
		}
		state = r.URL.Query()["state"]
		for _, e := range state {
			if !(e == "resolved" || e == "terminated") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("state[*]", e, []interface{}{"resolved", "terminated"}))
			}
		}
		appliedPolicyIDRaw := r.URL.Query().Get("applied_policy_id")
		if appliedPolicyIDRaw != "" {
			appliedPolicyID = &appliedPolicyIDRaw
		}
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		etagRaw := r.Header.Get("If-None-Match")
		if etagRaw != "" {
			etag = &etagRaw
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewIndexPayload(projectID, view, state, appliedPolicyID, apiVersion, etag, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeIndexError returns an encoder for errors returned by the index
// ArchivedIncident endpoint.
func EncodeIndexError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeIndexEscalationsResponse returns an encoder for responses returned by
// the ArchivedIncident index_escalations endpoint.
func EncodeIndexEscalationsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*archivedincidentviews.Escalations)
		enc := encoder(ctx, w)
		body := NewIndexEscalationsResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeIndexEscalationsRequest returns a decoder for requests sent to the
// ArchivedIncident index_escalations endpoint.
func DecodeIndexEscalationsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			incidentID string
			apiVersion string
			token      *string
			err        error

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		incidentID = params["incident_id"]
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewIndexEscalationsPayload(projectID, incidentID, apiVersion, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeIndexEscalationsError returns an encoder for errors returned by the
// index_escalations ArchivedIncident endpoint.
func EncodeIndexEscalationsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexEscalationsNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexEscalationsUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexEscalationsForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexEscalationsBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexEscalationsBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexEscalationsInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeIndexResolutionsResponse returns an encoder for responses returned by
// the ArchivedIncident index_resolutions endpoint.
func EncodeIndexResolutionsResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*archivedincidentviews.Resolutions)
		enc := encoder(ctx, w)
		body := NewIndexResolutionsResponseBody(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeIndexResolutionsRequest returns a decoder for requests sent to the
// ArchivedIncident index_resolutions endpoint.
func DecodeIndexResolutionsRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			projectID  uint
			incidentID string
			apiVersion string
			token      *string
			err        error

			params = mux.Vars(r)
		)
		{
			projectIDRaw := params["project_id"]
			v, err2 := strconv.ParseUint(projectIDRaw, 10, strconv.IntSize)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("projectID", projectIDRaw, "unsigned integer"))
			}
			projectID = uint(v)
		}
		incidentID = params["incident_id"]
		apiVersion = r.Header.Get("Api-Version")
		if apiVersion == "" {
			err = goa.MergeErrors(err, goa.MissingFieldError("Api-Version", "header"))
		}
		if !(apiVersion == "1.0") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("apiVersion", apiVersion, []interface{}{"1.0"}))
		}
		tokenRaw := r.Header.Get("Authorization")
		if tokenRaw != "" {
			token = &tokenRaw
		}
		if err != nil {
			return nil, err
		}
		payload := NewIndexResolutionsPayload(projectID, incidentID, apiVersion, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}

		return payload, nil
	}
}

// EncodeIndexResolutionsError returns an encoder for errors returned by the
// index_resolutions ArchivedIncident endpoint.
func EncodeIndexResolutionsError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexResolutionsNotFoundResponseBody(res)
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		case "unprocessable_entity":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexResolutionsUnprocessableEntityResponseBody(res)
			w.Header().Set("goa-error", "unprocessable_entity")
			w.WriteHeader(http.StatusUnprocessableEntity)
			return enc.Encode(body)
		case "unauthorized":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexResolutionsUnauthorizedResponseBody(res)
			w.Header().Set("goa-error", "unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
			return enc.Encode(body)
		case "forbidden":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexResolutionsForbiddenResponseBody(res)
			w.Header().Set("goa-error", "forbidden")
			w.WriteHeader(http.StatusForbidden)
			return enc.Encode(body)
		case "bad_request":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexResolutionsBadRequestResponseBody(res)
			w.Header().Set("goa-error", "bad_request")
			w.WriteHeader(http.StatusBadRequest)
			return enc.Encode(body)
		case "bad_gateway":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexResolutionsBadGatewayResponseBody(res)
			w.Header().Set("goa-error", "bad_gateway")
			w.WriteHeader(http.StatusBadGateway)
			return enc.Encode(body)
		case "internal_error":
			res := v.(*goa.ServiceError)
			enc := encoder(ctx, w)
			body := NewIndexResolutionsInternalErrorResponseBody(res)
			w.Header().Set("goa-error", "internal_error")
			w.WriteHeader(http.StatusInternalServerError)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// marshalArchivedincidentviewsProjectViewToProjectResponseBody builds a value
// of type *ProjectResponseBody from a value of type
// *archivedincidentviews.ProjectView.
func marshalArchivedincidentviewsProjectViewToProjectResponseBody(v *archivedincidentviews.ProjectView) *ProjectResponseBody {
	if v == nil {
		return nil
	}
	res := &ProjectResponseBody{
		ID:      *v.ID,
		Name:    *v.Name,
		OrgID:   *v.OrgID,
		OrgName: *v.OrgName,
	}

	return res
}

// marshalArchivedincidentviewsAppliedPolicyViewToAppliedPolicyResponseBodyLink
// builds a value of type *AppliedPolicyResponseBodyLink from a value of type
// *archivedincidentviews.AppliedPolicyView.
func marshalArchivedincidentviewsAppliedPolicyViewToAppliedPolicyResponseBodyLink(v *archivedincidentviews.AppliedPolicyView) *AppliedPolicyResponseBodyLink {
	if v == nil {
		return nil
	}
	res := &AppliedPolicyResponseBodyLink{
		ID:                  *v.ID,
		PolicyAggregateID:   v.PolicyAggregateID,
		IncidentAggregateID: v.IncidentAggregateID,
		Name:                *v.Name,
		Href:                *v.Href,
		CreatedAt:           v.CreatedAt,
		Frequency:           v.Frequency,
		Kind:                *v.Kind,
	}
	if v.PolicyTemplate != nil {
		res.PolicyTemplate = marshalArchivedincidentviewsPolicyTemplateViewToPolicyTemplateResponseBodyLink(v.PolicyTemplate)
	}
	if v.PublishedTemplate != nil {
		res.PublishedTemplate = marshalArchivedincidentviewsPublishedTemplateViewToPublishedTemplateResponseBodyLink(v.PublishedTemplate)
	}
	if v.CreatedBy != nil {
		res.CreatedBy = marshalArchivedincidentviewsUserViewToUserResponseBody(v.CreatedBy)
	}

	return res
}

// marshalArchivedincidentviewsPolicyTemplateViewToPolicyTemplateResponseBodyLink
// builds a value of type *PolicyTemplateResponseBodyLink from a value of type
// *archivedincidentviews.PolicyTemplateView.
func marshalArchivedincidentviewsPolicyTemplateViewToPolicyTemplateResponseBodyLink(v *archivedincidentviews.PolicyTemplateView) *PolicyTemplateResponseBodyLink {
	if v == nil {
		return nil
	}
	res := &PolicyTemplateResponseBodyLink{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Fingerprint: *v.Fingerprint,
		UpdatedAt:   v.UpdatedAt,
		Kind:        *v.Kind,
	}
	if v.UpdatedBy != nil {
		res.UpdatedBy = marshalArchivedincidentviewsUserViewToUserResponseBody(v.UpdatedBy)
	}

	return res
}

// marshalArchivedincidentviewsUserViewToUserResponseBody builds a value of
// type *UserResponseBody from a value of type *archivedincidentviews.UserView.
func marshalArchivedincidentviewsUserViewToUserResponseBody(v *archivedincidentviews.UserView) *UserResponseBody {
	if v == nil {
		return nil
	}
	res := &UserResponseBody{
		ID:    *v.ID,
		Email: *v.Email,
		Name:  *v.Name,
	}

	return res
}

// marshalArchivedincidentviewsPublishedTemplateViewToPublishedTemplateResponseBodyLink
// builds a value of type *PublishedTemplateResponseBodyLink from a value of
// type *archivedincidentviews.PublishedTemplateView.
func marshalArchivedincidentviewsPublishedTemplateViewToPublishedTemplateResponseBodyLink(v *archivedincidentviews.PublishedTemplateView) *PublishedTemplateResponseBodyLink {
	if v == nil {
		return nil
	}
	res := &PublishedTemplateResponseBodyLink{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Fingerprint: *v.Fingerprint,
		UpdatedAt:   v.UpdatedAt,
		BuiltIn:     v.BuiltIn,
		Kind:        *v.Kind,
	}
	if v.UpdatedBy != nil {
		res.UpdatedBy = marshalArchivedincidentviewsUserViewToUserResponseBody(v.UpdatedBy)
	}

	return res
}

// marshalArchivedincidentviewsConfigurationOptionViewToConfigurationOptionResponseBody
// builds a value of type *ConfigurationOptionResponseBody from a value of type
// *archivedincidentviews.ConfigurationOptionView.
func marshalArchivedincidentviewsConfigurationOptionViewToConfigurationOptionResponseBody(v *archivedincidentviews.ConfigurationOptionView) *ConfigurationOptionResponseBody {
	if v == nil {
		return nil
	}
	res := &ConfigurationOptionResponseBody{
		Name:  *v.Name,
		Label: *v.Label,
		Type:  *v.Type,
		Value: v.Value,
	}

	return res
}

// marshalArchivedincidentviewsAppliedPolicyViewToAppliedPolicyResponseBodySource
// builds a value of type *AppliedPolicyResponseBodySource from a value of type
// *archivedincidentviews.AppliedPolicyView.
func marshalArchivedincidentviewsAppliedPolicyViewToAppliedPolicyResponseBodySource(v *archivedincidentviews.AppliedPolicyView) *AppliedPolicyResponseBodySource {
	if v == nil {
		return nil
	}
	res := &AppliedPolicyResponseBodySource{
		ID:                  *v.ID,
		PolicyAggregateID:   v.PolicyAggregateID,
		IncidentAggregateID: v.IncidentAggregateID,
		Name:                *v.Name,
		Href:                *v.Href,
		Kind:                *v.Kind,
	}
	if v.Project != nil {
		res.Project = marshalArchivedincidentviewsProjectViewToProjectResponseBody(v.Project)
	}
	if v.PolicyTemplate != nil {
		res.PolicyTemplate = marshalArchivedincidentviewsPolicyTemplateViewToPolicyTemplateResponseBodySource(v.PolicyTemplate)
	}
	if v.PublishedTemplate != nil {
		res.PublishedTemplate = marshalArchivedincidentviewsPublishedTemplateViewToPublishedTemplateResponseBodySource(v.PublishedTemplate)
	}

	return res
}

// marshalArchivedincidentviewsPolicyTemplateViewToPolicyTemplateResponseBodySource
// builds a value of type *PolicyTemplateResponseBodySource from a value of
// type *archivedincidentviews.PolicyTemplateView.
func marshalArchivedincidentviewsPolicyTemplateViewToPolicyTemplateResponseBodySource(v *archivedincidentviews.PolicyTemplateView) *PolicyTemplateResponseBodySource {
	if v == nil {
		return nil
	}
	res := &PolicyTemplateResponseBodySource{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Filename:    v.Filename,
		Source:      v.Source,
		Fingerprint: *v.Fingerprint,
		Kind:        *v.Kind,
	}

	return res
}

// marshalArchivedincidentviewsPublishedTemplateViewToPublishedTemplateResponseBodySource
// builds a value of type *PublishedTemplateResponseBodySource from a value of
// type *archivedincidentviews.PublishedTemplateView.
func marshalArchivedincidentviewsPublishedTemplateViewToPublishedTemplateResponseBodySource(v *archivedincidentviews.PublishedTemplateView) *PublishedTemplateResponseBodySource {
	if v == nil {
		return nil
	}
	res := &PublishedTemplateResponseBodySource{
		ID:          *v.ID,
		Name:        *v.Name,
		Href:        *v.Href,
		Filename:    v.Filename,
		Source:      v.Source,
		Fingerprint: *v.Fingerprint,
		Kind:        *v.Kind,
	}

	return res
}

// marshalArchivedincidentviewsArchivedIncidentViewToArchivedIncidentResponseBody
// builds a value of type *ArchivedIncidentResponseBody from a value of type
// *archivedincidentviews.ArchivedIncidentView.
func marshalArchivedincidentviewsArchivedIncidentViewToArchivedIncidentResponseBody(v *archivedincidentviews.ArchivedIncidentView) *ArchivedIncidentResponseBody {
	if v == nil {
		return nil
	}
	res := &ArchivedIncidentResponseBody{
		ID:                 *v.ID,
		Href:               v.Href,
		Summary:            v.Summary,
		ViolationDataCount: v.ViolationDataCount,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
		ResolvedAt:         v.ResolvedAt,
		ResolutionMessage:  v.ResolutionMessage,
		State:              v.State,
		Severity:           v.Severity,
		Category:           v.Category,
		DryRun:             v.DryRun,
		ActionFailed:       v.ActionFailed,
		Kind:               *v.Kind,
		Etag:               *v.Etag,
		NotModified:        v.NotModified,
	}
	if v.Project != nil {
		res.Project = marshalArchivedincidentviewsProjectViewToProjectResponseBody(v.Project)
	}
	if v.AppliedPolicy != nil {
		res.AppliedPolicy = marshalArchivedincidentviewsAppliedPolicyViewToAppliedPolicyResponseBodyLink(v.AppliedPolicy)
	}
	if v.ResolvedBy != nil {
		res.ResolvedBy = marshalArchivedincidentviewsUserViewToUserResponseBody(v.ResolvedBy)
	}
	if v.Options != nil {
		res.Options = make([]*ConfigurationOptionResponseBody, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = marshalArchivedincidentviewsConfigurationOptionViewToConfigurationOptionResponseBody(val)
		}
	}

	return res
}

// marshalArchivedincidentviewsArchivedIncidentViewToArchivedIncidentResponseBodyExtended
// builds a value of type *ArchivedIncidentResponseBodyExtended from a value of
// type *archivedincidentviews.ArchivedIncidentView.
func marshalArchivedincidentviewsArchivedIncidentViewToArchivedIncidentResponseBodyExtended(v *archivedincidentviews.ArchivedIncidentView) *ArchivedIncidentResponseBodyExtended {
	if v == nil {
		return nil
	}
	res := &ArchivedIncidentResponseBodyExtended{
		ID:                 *v.ID,
		Href:               v.Href,
		Summary:            v.Summary,
		Detail:             v.Detail,
		ViolationDataCount: v.ViolationDataCount,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
		ResolvedAt:         v.ResolvedAt,
		ResolutionMessage:  v.ResolutionMessage,
		State:              v.State,
		Severity:           v.Severity,
		Category:           v.Category,
		DryRun:             v.DryRun,
		ActionFailed:       v.ActionFailed,
		Kind:               *v.Kind,
		Etag:               *v.Etag,
		NotModified:        v.NotModified,
	}
	if v.Project != nil {
		res.Project = marshalArchivedincidentviewsProjectViewToProjectResponseBody(v.Project)
	}
	if v.AppliedPolicy != nil {
		res.AppliedPolicy = marshalArchivedincidentviewsAppliedPolicyViewToAppliedPolicyResponseBodyLink(v.AppliedPolicy)
	}
	if v.ResolvedBy != nil {
		res.ResolvedBy = marshalArchivedincidentviewsUserViewToUserResponseBody(v.ResolvedBy)
	}
	if v.Options != nil {
		res.Options = make([]*ConfigurationOptionResponseBody, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = marshalArchivedincidentviewsConfigurationOptionViewToConfigurationOptionResponseBody(val)
		}
	}

	return res
}

// marshalArchivedincidentviewsEscalationViewToEscalationResponseBody builds a
// value of type *EscalationResponseBody from a value of type
// *archivedincidentviews.EscalationView.
func marshalArchivedincidentviewsEscalationViewToEscalationResponseBody(v *archivedincidentviews.EscalationView) *EscalationResponseBody {
	if v == nil {
		return nil
	}
	res := &EscalationResponseBody{
		Status: *v.Status,
		Name:   *v.Name,
	}
	if v.Actions != nil {
		res.Actions = make([]*EscalationActionResponseBody, len(v.Actions))
		for i, val := range v.Actions {
			res.Actions[i] = marshalArchivedincidentviewsEscalationActionViewToEscalationActionResponseBody(val)
		}
	}

	return res
}

// marshalArchivedincidentviewsEscalationActionViewToEscalationActionResponseBody
// builds a value of type *EscalationActionResponseBody from a value of type
// *archivedincidentviews.EscalationActionView.
func marshalArchivedincidentviewsEscalationActionViewToEscalationActionResponseBody(v *archivedincidentviews.EscalationActionView) *EscalationActionResponseBody {
	res := &EscalationActionResponseBody{
		Type:        *v.Type,
		Status:      *v.Status,
		StartedAt:   v.StartedAt,
		FinishedAt:  v.FinishedAt,
		Error:       v.Error,
		ProcessHref: v.ProcessHref,
	}
	if v.ApprovalRequest != nil {
		res.ApprovalRequest = marshalArchivedincidentviewsApprovalRequestViewToApprovalRequestResponseBodyExtended(v.ApprovalRequest)
	}

	return res
}

// marshalArchivedincidentviewsApprovalRequestViewToApprovalRequestResponseBodyExtended
// builds a value of type *ApprovalRequestResponseBodyExtended from a value of
// type *archivedincidentviews.ApprovalRequestView.
func marshalArchivedincidentviewsApprovalRequestViewToApprovalRequestResponseBodyExtended(v *archivedincidentviews.ApprovalRequestView) *ApprovalRequestResponseBodyExtended {
	if v == nil {
		return nil
	}
	res := &ApprovalRequestResponseBodyExtended{
		ID:            *v.ID,
		ProjectID:     *v.ProjectID,
		Href:          *v.Href,
		Label:         v.Label,
		Description:   v.Description,
		CreatedAt:     v.CreatedAt,
		UpdatedAt:     v.UpdatedAt,
		Status:        v.Status,
		ApprovedAt:    v.ApprovedAt,
		DenialComment: v.DenialComment,
		DeniedAt:      v.DeniedAt,
		Kind:          *v.Kind,
	}
	if v.Subject != nil {
		res.Subject = marshalArchivedincidentviewsApprovalSubjectToApprovalSubject(v.Subject)
	}
	if v.Parameters != nil {
		res.Parameters = make(map[string]*ParameterResponseBody, len(v.Parameters))
		for key, val := range v.Parameters {
			tk := key
			res.Parameters[tk] = marshalArchivedincidentviewsParameterViewToParameterResponseBody(val)
		}
	}
	if v.Options != nil {
		res.Options = make([]*ConfigurationOptionResponseBody, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = marshalArchivedincidentviewsConfigurationOptionViewToConfigurationOptionResponseBody(val)
		}
	}
	if v.ApprovedBy != nil {
		res.ApprovedBy = marshalArchivedincidentviewsUserViewToUserResponseBody(v.ApprovedBy)
	}
	if v.DeniedBy != nil {
		res.DeniedBy = marshalArchivedincidentviewsUserViewToUserResponseBody(v.DeniedBy)
	}

	return res
}

// marshalArchivedincidentviewsApprovalSubjectToApprovalSubject builds a value
// of type *ApprovalSubject from a value of type
// *archivedincidentviews.ApprovalSubject.
func marshalArchivedincidentviewsApprovalSubjectToApprovalSubject(v *archivedincidentviews.ApprovalSubject) *ApprovalSubject {
	res := &ApprovalSubject{
		Kind: *v.Kind,
		Href: *v.Href,
	}

	return res
}

// marshalArchivedincidentviewsParameterViewToParameterResponseBody builds a
// value of type *ParameterResponseBody from a value of type
// *archivedincidentviews.ParameterView.
func marshalArchivedincidentviewsParameterViewToParameterResponseBody(v *archivedincidentviews.ParameterView) *ParameterResponseBody {
	if v == nil {
		return nil
	}
	res := &ParameterResponseBody{
		Name:                  *v.Name,
		Type:                  *v.Type,
		Label:                 *v.Label,
		Index:                 *v.Index,
		Category:              v.Category,
		Description:           v.Description,
		Default:               v.Default,
		MinLength:             v.MinLength,
		MaxLength:             v.MaxLength,
		MinValue:              v.MinValue,
		MaxValue:              v.MaxValue,
		ConstraintDescription: v.ConstraintDescription,
	}
	if v.NoEcho != nil {
		res.NoEcho = *v.NoEcho
	}
	if v.NoEcho == nil {
		res.NoEcho = false
	}
	if v.AllowedValues != nil {
		res.AllowedValues = make([]interface{}, len(v.AllowedValues))
		for i, val := range v.AllowedValues {
			res.AllowedValues[i] = val
		}
	}
	if v.AllowedPattern != nil {
		res.AllowedPattern = marshalArchivedincidentviewsRegexpViewToRegexpResponseBody(v.AllowedPattern)
	}

	return res
}

// marshalArchivedincidentviewsRegexpViewToRegexpResponseBody builds a value of
// type *RegexpResponseBody from a value of type
// *archivedincidentviews.RegexpView.
func marshalArchivedincidentviewsRegexpViewToRegexpResponseBody(v *archivedincidentviews.RegexpView) *RegexpResponseBody {
	if v == nil {
		return nil
	}
	res := &RegexpResponseBody{
		Pattern: *v.Pattern,
		Options: v.Options,
	}

	return res
}

// marshalArchivedincidentviewsResolutionViewToResolutionResponseBody builds a
// value of type *ResolutionResponseBody from a value of type
// *archivedincidentviews.ResolutionView.
func marshalArchivedincidentviewsResolutionViewToResolutionResponseBody(v *archivedincidentviews.ResolutionView) *ResolutionResponseBody {
	if v == nil {
		return nil
	}
	res := &ResolutionResponseBody{
		Status: *v.Status,
		Name:   *v.Name,
	}
	if v.Actions != nil {
		res.Actions = make([]*ResolutionActionResponseBody, len(v.Actions))
		for i, val := range v.Actions {
			res.Actions[i] = marshalArchivedincidentviewsResolutionActionViewToResolutionActionResponseBody(val)
		}
	}

	return res
}

// marshalArchivedincidentviewsResolutionActionViewToResolutionActionResponseBody
// builds a value of type *ResolutionActionResponseBody from a value of type
// *archivedincidentviews.ResolutionActionView.
func marshalArchivedincidentviewsResolutionActionViewToResolutionActionResponseBody(v *archivedincidentviews.ResolutionActionView) *ResolutionActionResponseBody {
	res := &ResolutionActionResponseBody{
		Type:        *v.Type,
		Status:      *v.Status,
		StartedAt:   v.StartedAt,
		FinishedAt:  v.FinishedAt,
		Error:       v.Error,
		ProcessHref: v.ProcessHref,
	}
	if v.ApprovalRequest != nil {
		res.ApprovalRequest = marshalArchivedincidentviewsApprovalRequestViewToApprovalRequestResponseBodyExtended(v.ApprovalRequest)
	}

	return res
}
