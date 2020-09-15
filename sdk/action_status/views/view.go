// Code generated by goa v3.1.3, DO NOT EDIT.
//
// ActionStatus views
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// ActionStatusList is the viewed result type that is projected based on a view.
type ActionStatusList struct {
	// Type to project
	Projected *ActionStatusListView
	// View to render
	View string
}

// ActionStatus is the viewed result type that is projected based on a view.
type ActionStatus struct {
	// Type to project
	Projected *ActionStatusView
	// View to render
	View string
}

// ActionStatusListView is a type that runs validations on a projected type.
type ActionStatusListView struct {
	// count is the number of action statuss in the list.
	Count *uint
	// etag is an HTTP ETag for the action list.
	Etag *string
	// items is the array of action statuses.
	Items ActionStatusCollectionView
	// not_modified is a flag used internally that indicates how to encode the HTTP
	// response (i.e. 200 or 304).
	NotModified *string
	// kind is "gov#action_status_list".
	Kind *string
}

// ActionStatusCollectionView is a type that runs validations on a projected
// type.
type ActionStatusCollectionView []*ActionStatusView

// ActionStatusView is a type that runs validations on a projected type.
type ActionStatusView struct {
	// id is a unique identifier for this action status.
	ID *string
	// status of the action.
	Status *string
	// name is an identifier for the action.
	Name *string
	// type is the type of action, which also indicates the situations in which it
	// is used.
	Type *string
	// label is the human readable name for the action.
	Label *string
	// run_by is the Flexera user that ran the action. If the action was applied
	// automatically, this will be who applied the action.
	RunBy *UserView
	// whether or not this action is automatically run.
	Automatic *bool
	// options lists the configuration options used to parameterize the action.
	Options []*ConfigurationOptionView
	// started_at is the time when the action was started.
	StartedAt *string
	// finished_at is the time when the action was finished.
	FinishedAt *string
	// actions is the list of individual steps within this action and their success
	// or failure.
	ActionItems []*ActionItemStatusView
	// kind is "gov#action_status".
	Kind *string
}

// UserView is a type that runs validations on a projected type.
type UserView struct {
	// ID of user
	ID *uint
	// email of user
	Email *string
	// name of user, usually of the form 'First Last'
	Name *string
}

// ConfigurationOptionView is a type that runs validations on a projected type.
type ConfigurationOptionView struct {
	// name of option
	Name *string
	// label of option
	Label *string
	// type of option
	Type *string
	// value of option
	Value interface{}
	// no_echo determines whether the value of the configuration option should be
	// hidden in UIs and API responses.
	NoEcho *bool
}

// ActionItemStatusView is a type that runs validations on a projected type.
type ActionItemStatusView struct {
	// type of the action item.
	Type *string
	// status of the action item.
	Status *string
	// started_at is the time when the action item was started.
	StartedAt *string
	// finished_at is the time when the action item was finished.
	FinishedAt *string
	// error is any error that occurred when handling the action item.
	Error *string
	// approval_request_href is an href of the approval request. Required if the
	// type is request_approval.
	ApprovalRequestHref *string
	// process_href is a url of a cloud workflow process. Required if the type is
	// cloud_workflow.
	ProcessHref *string
}

var (
	// ActionStatusListMap is a map of attribute names in result type
	// ActionStatusList indexed by view name.
	ActionStatusListMap = map[string][]string{
		"default": []string{
			"count",
			"items",
			"etag",
			"not_modified",
			"kind",
		},
		"extended": []string{
			"count",
			"items",
			"etag",
			"not_modified",
			"kind",
		},
	}
	// ActionStatusMap is a map of attribute names in result type ActionStatus
	// indexed by view name.
	ActionStatusMap = map[string][]string{
		"default": []string{
			"id",
			"status",
			"type",
			"automatic",
			"name",
			"label",
			"run_by",
			"started_at",
			"options",
			"finished_at",
			"kind",
		},
		"extended": []string{
			"id",
			"status",
			"type",
			"automatic",
			"name",
			"label",
			"run_by",
			"started_at",
			"options",
			"finished_at",
			"kind",
			"action_items",
		},
	}
	// ActionStatusCollectionMap is a map of attribute names in result type
	// ActionStatusCollection indexed by view name.
	ActionStatusCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"status",
			"type",
			"automatic",
			"name",
			"label",
			"run_by",
			"started_at",
			"options",
			"finished_at",
			"kind",
		},
		"extended": []string{
			"id",
			"status",
			"type",
			"automatic",
			"name",
			"label",
			"run_by",
			"started_at",
			"options",
			"finished_at",
			"kind",
			"action_items",
		},
	}
	// ActionItemStatusMap is a map of attribute names in result type
	// ActionItemStatus indexed by view name.
	ActionItemStatusMap = map[string][]string{
		"default": []string{
			"type",
			"status",
			"started_at",
			"finished_at",
			"error",
			"approval_request_href",
			"process_href",
		},
	}
)

// ValidateActionStatusList runs the validations defined on the viewed result
// type ActionStatusList.
func ValidateActionStatusList(result *ActionStatusList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateActionStatusListView(result.Projected)
	case "extended":
		err = ValidateActionStatusListViewExtended(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "extended"})
	}
	return
}

// ValidateActionStatus runs the validations defined on the viewed result type
// ActionStatus.
func ValidateActionStatus(result *ActionStatus) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateActionStatusView(result.Projected)
	case "extended":
		err = ValidateActionStatusViewExtended(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "extended"})
	}
	return
}

// ValidateActionStatusListView runs the validations defined on
// ActionStatusListView using the "default" view.
func ValidateActionStatusListView(result *ActionStatusListView) (err error) {
	if result.Etag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("etag", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.NotModified != nil {
		if !(*result.NotModified == "true" || *result.NotModified == "false") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.not_modified", *result.NotModified, []interface{}{"true", "false"}))
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#action_status_list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#action_status_list"}))
		}
	}
	if result.Items != nil {
		if err2 := ValidateActionStatusCollectionView(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActionStatusListViewExtended runs the validations defined on
// ActionStatusListView using the "extended" view.
func ValidateActionStatusListViewExtended(result *ActionStatusListView) (err error) {
	if result.Etag == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("etag", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.NotModified != nil {
		if !(*result.NotModified == "true" || *result.NotModified == "false") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.not_modified", *result.NotModified, []interface{}{"true", "false"}))
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#action_status_list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#action_status_list"}))
		}
	}
	if result.Items != nil {
		if err2 := ValidateActionStatusCollectionViewExtended(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActionStatusCollectionView runs the validations defined on
// ActionStatusCollectionView using the "default" view.
func ValidateActionStatusCollectionView(result ActionStatusCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateActionStatusView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActionStatusCollectionViewExtended runs the validations defined on
// ActionStatusCollectionView using the "extended" view.
func ValidateActionStatusCollectionViewExtended(result ActionStatusCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateActionStatusViewExtended(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateActionStatusView runs the validations defined on ActionStatusView
// using the "default" view.
func ValidateActionStatusView(result *ActionStatusView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "queued" || *result.Status == "aborted" || *result.Status == "pending" || *result.Status == "running" || *result.Status == "completed" || *result.Status == "failed" || *result.Status == "denied") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"queued", "aborted", "pending", "running", "completed", "failed", "denied"}))
		}
	}
	if result.Type != nil {
		if !(*result.Type == "escalation" || *result.Type == "resolution") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"escalation", "resolution"}))
		}
	}
	if result.RunBy != nil {
		if err2 := ValidateUserView(result.RunBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	for _, e := range result.Options {
		if e != nil {
			if err2 := ValidateConfigurationOptionView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.finished_at", *result.FinishedAt, goa.FormatDateTime))
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#action_status") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#action_status"}))
		}
	}
	return
}

// ValidateActionStatusViewExtended runs the validations defined on
// ActionStatusView using the "extended" view.
func ValidateActionStatusViewExtended(result *ActionStatusView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.ActionItems == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("action_items", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Status != nil {
		if !(*result.Status == "queued" || *result.Status == "aborted" || *result.Status == "pending" || *result.Status == "running" || *result.Status == "completed" || *result.Status == "failed" || *result.Status == "denied") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"queued", "aborted", "pending", "running", "completed", "failed", "denied"}))
		}
	}
	if result.Type != nil {
		if !(*result.Type == "escalation" || *result.Type == "resolution") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"escalation", "resolution"}))
		}
	}
	if result.RunBy != nil {
		if err2 := ValidateUserView(result.RunBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	for _, e := range result.Options {
		if e != nil {
			if err2 := ValidateConfigurationOptionView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.finished_at", *result.FinishedAt, goa.FormatDateTime))
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#action_status") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#action_status"}))
		}
	}
	for _, e := range result.ActionItems {
		if e != nil {
			if err2 := ValidateActionItemStatusView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateUserView runs the validations defined on UserView.
func ValidateUserView(result *UserView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.email", *result.Email, goa.FormatEmail))
	}
	return
}

// ValidateConfigurationOptionView runs the validations defined on
// ConfigurationOptionView.
func ValidateConfigurationOptionView(result *ConfigurationOptionView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "string" || *result.Type == "number" || *result.Type == "list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"string", "number", "list"}))
		}
	}
	return
}

// ValidateActionItemStatusView runs the validations defined on
// ActionItemStatusView using the "default" view.
func ValidateActionItemStatusView(result *ActionItemStatusView) (err error) {
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Status == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("status", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "email" || *result.Type == "cloud_workflow" || *result.Type == "request_approval" || *result.Type == "resolve_incident") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"email", "cloud_workflow", "request_approval", "resolve_incident"}))
		}
	}
	if result.Status != nil {
		if !(*result.Status == "queued" || *result.Status == "aborted" || *result.Status == "pending" || *result.Status == "running" || *result.Status == "completed" || *result.Status == "skipped" || *result.Status == "failed" || *result.Status == "approved" || *result.Status == "denied") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.status", *result.Status, []interface{}{"queued", "aborted", "pending", "running", "completed", "skipped", "failed", "approved", "denied"}))
		}
	}
	if result.StartedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.started_at", *result.StartedAt, goa.FormatDateTime))
	}
	if result.FinishedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.finished_at", *result.FinishedAt, goa.FormatDateTime))
	}
	if result.ApprovalRequestHref != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.approval_request_href", *result.ApprovalRequestHref, "^/api/governance/projects/[0-9]+/approval_requests/[0-9a-f]+$"))
	}
	return
}