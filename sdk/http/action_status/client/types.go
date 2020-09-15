// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ActionStatus HTTP client types
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package client

import (
	actionstatusviews "github.com/rightscale/policy_sdk/sdk/action_status/views"
	goa "goa.design/goa/v3/pkg"
)

// IndexOKResponseBody is the type of the "ActionStatus" service "index"
// endpoint HTTP response body.
type IndexOKResponseBody struct {
	// count is the number of action statuss in the list.
	Count *uint `form:"count,omitempty" json:"count,omitempty" xml:"count,omitempty"`
	// items is the array of action statuses.
	Items ActionStatusCollectionResponseBody `form:"items,omitempty" json:"items,omitempty" xml:"items,omitempty"`
	// not_modified is a flag used internally that indicates how to encode the HTTP
	// response (i.e. 200 or 304).
	NotModified *string `form:"not_modified,omitempty" json:"not_modified,omitempty" xml:"not_modified,omitempty"`
	// kind is "gov#action_status_list".
	Kind *string `form:"kind,omitempty" json:"kind,omitempty" xml:"kind,omitempty"`
}

// ShowResponseBody is the type of the "ActionStatus" service "show" endpoint
// HTTP response body.
type ShowResponseBody struct {
	// id is a unique identifier for this action status.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// status of the action.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// name is an identifier for the action.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// type is the type of action, which also indicates the situations in which it
	// is used.
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// label is the human readable name for the action.
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// run_by is the Flexera user that ran the action. If the action was applied
	// automatically, this will be who applied the action.
	RunBy *UserResponseBody `form:"run_by,omitempty" json:"run_by,omitempty" xml:"run_by,omitempty"`
	// whether or not this action is automatically run.
	Automatic *bool `form:"automatic,omitempty" json:"automatic,omitempty" xml:"automatic,omitempty"`
	// options lists the configuration options used to parameterize the action.
	Options []*ConfigurationOptionResponseBody `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	// started_at is the time when the action was started.
	StartedAt *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	// finished_at is the time when the action was finished.
	FinishedAt *string `form:"finished_at,omitempty" json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// actions is the list of individual steps within this action and their success
	// or failure.
	ActionItems []*ActionItemStatusResponseBody `form:"action_items,omitempty" json:"action_items,omitempty" xml:"action_items,omitempty"`
	// kind is "gov#action_status".
	Kind *string `form:"kind,omitempty" json:"kind,omitempty" xml:"kind,omitempty"`
}

// IndexUnauthorizedResponseBody is the type of the "ActionStatus" service
// "index" endpoint HTTP response body for the "unauthorized" error.
type IndexUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// IndexForbiddenResponseBody is the type of the "ActionStatus" service "index"
// endpoint HTTP response body for the "forbidden" error.
type IndexForbiddenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// IndexBadRequestResponseBody is the type of the "ActionStatus" service
// "index" endpoint HTTP response body for the "bad_request" error.
type IndexBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// IndexBadGatewayResponseBody is the type of the "ActionStatus" service
// "index" endpoint HTTP response body for the "bad_gateway" error.
type IndexBadGatewayResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// IndexInternalErrorResponseBody is the type of the "ActionStatus" service
// "index" endpoint HTTP response body for the "internal_error" error.
type IndexInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ShowNotFoundResponseBody is the type of the "ActionStatus" service "show"
// endpoint HTTP response body for the "not_found" error.
type ShowNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ShowUnauthorizedResponseBody is the type of the "ActionStatus" service
// "show" endpoint HTTP response body for the "unauthorized" error.
type ShowUnauthorizedResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ShowForbiddenResponseBody is the type of the "ActionStatus" service "show"
// endpoint HTTP response body for the "forbidden" error.
type ShowForbiddenResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ShowBadRequestResponseBody is the type of the "ActionStatus" service "show"
// endpoint HTTP response body for the "bad_request" error.
type ShowBadRequestResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ShowBadGatewayResponseBody is the type of the "ActionStatus" service "show"
// endpoint HTTP response body for the "bad_gateway" error.
type ShowBadGatewayResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ShowInternalErrorResponseBody is the type of the "ActionStatus" service
// "show" endpoint HTTP response body for the "internal_error" error.
type ShowInternalErrorResponseBody struct {
	// Name is the name of this class of errors.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message *string `form:"message,omitempty" json:"message,omitempty" xml:"message,omitempty"`
	// Is the error temporary?
	Temporary *bool `form:"temporary,omitempty" json:"temporary,omitempty" xml:"temporary,omitempty"`
	// Is the error a timeout?
	Timeout *bool `form:"timeout,omitempty" json:"timeout,omitempty" xml:"timeout,omitempty"`
	// Is the error a server-side fault?
	Fault *bool `form:"fault,omitempty" json:"fault,omitempty" xml:"fault,omitempty"`
}

// ActionStatusCollectionResponseBody is used to define fields on response body
// types.
type ActionStatusCollectionResponseBody []*ActionStatusResponseBody

// ActionStatusResponseBody is used to define fields on response body types.
type ActionStatusResponseBody struct {
	// id is a unique identifier for this action status.
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// status of the action.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// name is an identifier for the action.
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// type is the type of action, which also indicates the situations in which it
	// is used.
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// label is the human readable name for the action.
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// run_by is the Flexera user that ran the action. If the action was applied
	// automatically, this will be who applied the action.
	RunBy *UserResponseBody `form:"run_by,omitempty" json:"run_by,omitempty" xml:"run_by,omitempty"`
	// whether or not this action is automatically run.
	Automatic *bool `form:"automatic,omitempty" json:"automatic,omitempty" xml:"automatic,omitempty"`
	// options lists the configuration options used to parameterize the action.
	Options []*ConfigurationOptionResponseBody `form:"options,omitempty" json:"options,omitempty" xml:"options,omitempty"`
	// started_at is the time when the action was started.
	StartedAt *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	// finished_at is the time when the action was finished.
	FinishedAt *string `form:"finished_at,omitempty" json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// actions is the list of individual steps within this action and their success
	// or failure.
	ActionItems []*ActionItemStatusResponseBody `form:"action_items,omitempty" json:"action_items,omitempty" xml:"action_items,omitempty"`
	// kind is "gov#action_status".
	Kind *string `form:"kind,omitempty" json:"kind,omitempty" xml:"kind,omitempty"`
}

// UserResponseBody is used to define fields on response body types.
type UserResponseBody struct {
	// ID of user
	ID *uint `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// email of user
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// name of user, usually of the form 'First Last'
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// ConfigurationOptionResponseBody is used to define fields on response body
// types.
type ConfigurationOptionResponseBody struct {
	// name of option
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// label of option
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
	// type of option
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// value of option
	Value interface{} `form:"value,omitempty" json:"value,omitempty" xml:"value,omitempty"`
	// no_echo determines whether the value of the configuration option should be
	// hidden in UIs and API responses.
	NoEcho *bool `form:"no_echo,omitempty" json:"no_echo,omitempty" xml:"no_echo,omitempty"`
}

// ActionItemStatusResponseBody is used to define fields on response body types.
type ActionItemStatusResponseBody struct {
	// type of the action item.
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// status of the action item.
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// started_at is the time when the action item was started.
	StartedAt *string `form:"started_at,omitempty" json:"started_at,omitempty" xml:"started_at,omitempty"`
	// finished_at is the time when the action item was finished.
	FinishedAt *string `form:"finished_at,omitempty" json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// error is any error that occurred when handling the action item.
	Error *string `form:"error,omitempty" json:"error,omitempty" xml:"error,omitempty"`
	// approval_request_href is an href of the approval request. Required if the
	// type is request_approval.
	ApprovalRequestHref *string `form:"approval_request_href,omitempty" json:"approval_request_href,omitempty" xml:"approval_request_href,omitempty"`
	// process_href is a url of a cloud workflow process. Required if the type is
	// cloud_workflow.
	ProcessHref *string `form:"process_href,omitempty" json:"process_href,omitempty" xml:"process_href,omitempty"`
}

// NewIndexActionStatusListNotModified builds a "ActionStatus" service "index"
// endpoint result from a HTTP "NotModified" response.
func NewIndexActionStatusListNotModified(etag string) *actionstatusviews.ActionStatusListView {
	v := &actionstatusviews.ActionStatusListView{}
	v.Etag = &etag

	return v
}

// NewIndexActionStatusListOK builds a "ActionStatus" service "index" endpoint
// result from a HTTP "OK" response.
func NewIndexActionStatusListOK(body *IndexOKResponseBody, etag string) *actionstatusviews.ActionStatusListView {
	v := &actionstatusviews.ActionStatusListView{
		Count:       body.Count,
		NotModified: body.NotModified,
		Kind:        body.Kind,
	}
	if body.Items != nil {
		v.Items = make([]*actionstatusviews.ActionStatusView, len(body.Items))
		for i, val := range body.Items {
			v.Items[i] = unmarshalActionStatusResponseBodyToActionstatusviewsActionStatusView(val)
		}
	}
	v.Etag = &etag

	return v
}

// NewIndexUnauthorized builds a ActionStatus service index endpoint
// unauthorized error.
func NewIndexUnauthorized(body *IndexUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewIndexForbidden builds a ActionStatus service index endpoint forbidden
// error.
func NewIndexForbidden(body *IndexForbiddenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewIndexBadRequest builds a ActionStatus service index endpoint bad_request
// error.
func NewIndexBadRequest(body *IndexBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewIndexBadGateway builds a ActionStatus service index endpoint bad_gateway
// error.
func NewIndexBadGateway(body *IndexBadGatewayResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewIndexInternalError builds a ActionStatus service index endpoint
// internal_error error.
func NewIndexInternalError(body *IndexInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewShowActionStatusOK builds a "ActionStatus" service "show" endpoint result
// from a HTTP "OK" response.
func NewShowActionStatusOK(body *ShowResponseBody) *actionstatusviews.ActionStatusView {
	v := &actionstatusviews.ActionStatusView{
		ID:         body.ID,
		Status:     body.Status,
		Name:       body.Name,
		Type:       body.Type,
		Label:      body.Label,
		Automatic:  body.Automatic,
		StartedAt:  body.StartedAt,
		FinishedAt: body.FinishedAt,
		Kind:       body.Kind,
	}
	if body.RunBy != nil {
		v.RunBy = unmarshalUserResponseBodyToActionstatusviewsUserView(body.RunBy)
	}
	if body.Options != nil {
		v.Options = make([]*actionstatusviews.ConfigurationOptionView, len(body.Options))
		for i, val := range body.Options {
			v.Options[i] = unmarshalConfigurationOptionResponseBodyToActionstatusviewsConfigurationOptionView(val)
		}
	}
	v.ActionItems = make([]*actionstatusviews.ActionItemStatusView, len(body.ActionItems))
	for i, val := range body.ActionItems {
		v.ActionItems[i] = unmarshalActionItemStatusResponseBodyToActionstatusviewsActionItemStatusView(val)
	}

	return v
}

// NewShowNotFound builds a ActionStatus service show endpoint not_found error.
func NewShowNotFound(body *ShowNotFoundResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewShowUnauthorized builds a ActionStatus service show endpoint unauthorized
// error.
func NewShowUnauthorized(body *ShowUnauthorizedResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewShowForbidden builds a ActionStatus service show endpoint forbidden error.
func NewShowForbidden(body *ShowForbiddenResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewShowBadRequest builds a ActionStatus service show endpoint bad_request
// error.
func NewShowBadRequest(body *ShowBadRequestResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewShowBadGateway builds a ActionStatus service show endpoint bad_gateway
// error.
func NewShowBadGateway(body *ShowBadGatewayResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// NewShowInternalError builds a ActionStatus service show endpoint
// internal_error error.
func NewShowInternalError(body *ShowInternalErrorResponseBody) *goa.ServiceError {
	v := &goa.ServiceError{
		Name:      *body.Name,
		ID:        *body.ID,
		Message:   *body.Message,
		Temporary: *body.Temporary,
		Timeout:   *body.Timeout,
		Fault:     *body.Fault,
	}

	return v
}

// ValidateIndexUnauthorizedResponseBody runs the validations defined on
// index_unauthorized_response_body
func ValidateIndexUnauthorizedResponseBody(body *IndexUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateIndexForbiddenResponseBody runs the validations defined on
// index_forbidden_response_body
func ValidateIndexForbiddenResponseBody(body *IndexForbiddenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateIndexBadRequestResponseBody runs the validations defined on
// index_bad_request_response_body
func ValidateIndexBadRequestResponseBody(body *IndexBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateIndexBadGatewayResponseBody runs the validations defined on
// index_bad_gateway_response_body
func ValidateIndexBadGatewayResponseBody(body *IndexBadGatewayResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateIndexInternalErrorResponseBody runs the validations defined on
// index_internal_error_response_body
func ValidateIndexInternalErrorResponseBody(body *IndexInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateShowNotFoundResponseBody runs the validations defined on
// show_not_found_response_body
func ValidateShowNotFoundResponseBody(body *ShowNotFoundResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateShowUnauthorizedResponseBody runs the validations defined on
// show_unauthorized_response_body
func ValidateShowUnauthorizedResponseBody(body *ShowUnauthorizedResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateShowForbiddenResponseBody runs the validations defined on
// show_forbidden_response_body
func ValidateShowForbiddenResponseBody(body *ShowForbiddenResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateShowBadRequestResponseBody runs the validations defined on
// show_bad_request_response_body
func ValidateShowBadRequestResponseBody(body *ShowBadRequestResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateShowBadGatewayResponseBody runs the validations defined on
// show_bad_gateway_response_body
func ValidateShowBadGatewayResponseBody(body *ShowBadGatewayResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateShowInternalErrorResponseBody runs the validations defined on
// show_internal_error_response_body
func ValidateShowInternalErrorResponseBody(body *ShowInternalErrorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Message == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("message", "body"))
	}
	if body.Temporary == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("temporary", "body"))
	}
	if body.Timeout == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("timeout", "body"))
	}
	if body.Fault == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fault", "body"))
	}
	return
}

// ValidateActionStatusCollectionResponseBody runs the validations defined on
// ActionStatusCollectionResponseBody
func ValidateActionStatusCollectionResponseBody(body ActionStatusCollectionResponseBody) (err error) {
	for _, e := range body {
		if e != nil {
			if err2 := ValidateActionStatusResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateActionStatusResponseBody runs the validations defined on
// ActionStatusResponseBody
func ValidateActionStatusResponseBody(body *ActionStatusResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.ActionItems == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action_items", "body"))
	}
	if body.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Status != nil {
		if !(*body.Status == "queued" || *body.Status == "aborted" || *body.Status == "pending" || *body.Status == "running" || *body.Status == "completed" || *body.Status == "failed" || *body.Status == "denied") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"queued", "aborted", "pending", "running", "completed", "failed", "denied"}))
		}
	}
	if body.Type != nil {
		if !(*body.Type == "escalation" || *body.Type == "resolution") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.type", *body.Type, []interface{}{"escalation", "resolution"}))
		}
	}
	if body.RunBy != nil {
		if err2 := ValidateUserResponseBody(body.RunBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range body.Options {
		if e != nil {
			if err2 := ValidateConfigurationOptionResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.started_at", *body.StartedAt, goa.FormatDateTime))
	}
	if body.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.finished_at", *body.FinishedAt, goa.FormatDateTime))
	}
	for _, e := range body.ActionItems {
		if e != nil {
			if err2 := ValidateActionItemStatusResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if body.Kind != nil {
		if !(*body.Kind == "gov#action_status") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.kind", *body.Kind, []interface{}{"gov#action_status"}))
		}
	}
	return
}

// ValidateUserResponseBody runs the validations defined on UserResponseBody
func ValidateUserResponseBody(body *UserResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	return
}

// ValidateConfigurationOptionResponseBody runs the validations defined on
// ConfigurationOptionResponseBody
func ValidateConfigurationOptionResponseBody(body *ConfigurationOptionResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "body"))
	}
	if body.Type != nil {
		if !(*body.Type == "string" || *body.Type == "number" || *body.Type == "list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.type", *body.Type, []interface{}{"string", "number", "list"}))
		}
	}
	return
}

// ValidateActionItemStatusResponseBody runs the validations defined on
// ActionItemStatusResponseBody
func ValidateActionItemStatusResponseBody(body *ActionItemStatusResponseBody) (err error) {
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "body"))
	}
	if body.Type != nil {
		if !(*body.Type == "email" || *body.Type == "cloud_workflow" || *body.Type == "request_approval" || *body.Type == "resolve_incident") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.type", *body.Type, []interface{}{"email", "cloud_workflow", "request_approval", "resolve_incident"}))
		}
	}
	if body.Status != nil {
		if !(*body.Status == "queued" || *body.Status == "aborted" || *body.Status == "pending" || *body.Status == "running" || *body.Status == "completed" || *body.Status == "skipped" || *body.Status == "failed" || *body.Status == "approved" || *body.Status == "denied") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.status", *body.Status, []interface{}{"queued", "aborted", "pending", "running", "completed", "skipped", "failed", "approved", "denied"}))
		}
	}
	if body.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.started_at", *body.StartedAt, goa.FormatDateTime))
	}
	if body.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.finished_at", *body.FinishedAt, goa.FormatDateTime))
	}
	if body.ApprovalRequestHref != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("body.approval_request_href", *body.ApprovalRequestHref, "^/api/governance/projects/[0-9]+/approval_requests/[0-9a-f]+$"))
	}
	return
}