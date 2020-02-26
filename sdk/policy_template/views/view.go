// Code generated by goa v3.0.6, DO NOT EDIT.
//
// PolicyTemplate views
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package views

import (
	goa "goa.design/goa/v3/pkg"
)

// PolicyTemplate is the viewed result type that is projected based on a view.
type PolicyTemplate struct {
	// Type to project
	Projected *PolicyTemplateView
	// View to render
	View string
}

// PolicyTemplateList is the viewed result type that is projected based on a
// view.
type PolicyTemplateList struct {
	// Type to project
	Projected *PolicyTemplateListView
	// View to render
	View string
}

// PolicyTemplateView is a type that runs validations on a projected type.
type PolicyTemplateView struct {
	// id identifies a policy template by ID.
	ID *string
	// name is the unique name of the policy template in the project.
	Name *string
	// project_id is the ID of the project that the policy template applies to.
	ProjectID *uint
	// rs_pt_ver is the policy engine version.
	RsPtVer *uint
	// short_description is the short description of the policy template.
	ShortDescription *string
	// long_description is the long description of the policy template. The content
	// can be markdown.
	LongDescription *string
	// doc_link is an HTTP URI to a page containing detailed documentation for the
	// policy.
	DocLink *string
	// info is an arbitrary set of key/value pairs that provide additional
	// information such as the policy author, support contact information, etc.
	Info map[string]string
	// default_frequency defines the interval the template will be run unless set
	// differently during application.
	DefaultFrequency *string
	// href is the href of the policy template.
	Href *string
	// filename is the name of the file that was uploaded to create the policy
	// template.
	Filename *string
	// source is the policy template source code.
	Source *string
	// fingerprint is a SHA created during compilation. It is used to determine if
	// the template is outdated.
	Fingerprint *string
	// category is the type categorization of the policy template.
	Category *string
	// created_by is the RightScale user that created the policy template.
	CreatedBy *UserView
	// created_at is the policy template creation timestamp in RFC3339 format.
	CreatedAt *string
	// updated_by is the RightScale user that updated the policy template.
	UpdatedBy *UserView
	// updated_at is the last update timestamp in RFC3339 format.
	UpdatedAt *string
	// permissions is a list of permissions required to run the policy.
	Permissions map[string]*PermissionView
	// required_roles is a list of governance roles, derived from permissions,
	// required to run the policy.
	RequiredRoles []string
	// parameters is a list of parameters required to apply the policy.
	Parameters map[string]*ParameterView
	// severity defines the severity level of incidents raised from this policy
	// template.
	Severity *string
	// tenancy indicates whether this template can be run across multiple projects
	// or is restricted to a single project.
	Tenancy *string
	// credentials is a list of authorization for external APIs needed to run the
	// policy.
	Credentials map[string]*CredentialsView
	// kind is "gov#policy_template".
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

// PermissionView is a type that runs validations on a projected type.
type PermissionView struct {
	// Name of a permission
	Name *string `json:"name"`
	// Label is used in the UI
	Label *string `json:"label"`
	// List of resource names the permission is applied to
	Resources []string `json:"resources"`
	// List of action names the permission is applied to
	Actions []string `json:"actions"`
}

// ParameterView is a type that runs validations on a projected type.
type ParameterView struct {
	// Name of the parameter
	Name *string `json:"name"`
	// Type of the parameter
	Type *string `json:"type"`
	// Label to show in the UI
	Label *string `json:"label"`
	// The index of this parameter in the list
	Index *uint `json:"index"`
	// The category used to group parameters
	Category *string `json:"category"`
	// Description of the parameter
	Description *string `json:"description"`
	// The default value for the parameter
	Default interface{} `json:"default"`
	// no_echo determines whether the value of the parameter should be hidden in
	// UIs and API responses.
	NoEcho *bool `json:"no_echo"`
	// List of values allowed for this parameter
	AllowedValues []interface{} `json:"allowed_values"`
	// The minimum length of a string parameter
	MinLength *uint `json:"min_length"`
	// The maximum length of a string parameter
	MaxLength *uint `json:"max_length"`
	// The minimum value of a number parmameter
	MinValue *float64 `json:"min_value"`
	// The maximum value of a number parameter
	MaxValue *float64 `json:"max_value"`
	// The regular expression pattern used to validate a string parameter
	AllowedPattern *RegexpView `json:"allowed_pattern"`
	// The description used for constraints
	ConstraintDescription *string `json:"constraint_description"`
}

// RegexpView is a type that runs validations on a projected type.
type RegexpView struct {
	// Pattern is the regular expression pattern.
	Pattern *string `json:"pattern"`
	// Options are the regular expression options. Options i (case insensitve) and
	// m (match over newlines) supported.
	Options *string `json:"options"`
}

// CredentialsView is a type that runs validations on a projected type.
type CredentialsView struct {
	// Name in policy template source code
	Name *string `json:"name"`
	// Schemes of credentials service resource supported.
	Schemes []string `json:"schemes"`
	// Label for the auth reference
	Label *string `json:"label"`
	// Description of what types of permissions need to be provided by auth.
	Description *string `json:"description"`
	// Tags is an optional filter to only show credentials matching a certain tag.
	Tags []*CredentialsTagView `json:"tags"`
}

// CredentialsTagView is a type that runs validations on a projected type.
type CredentialsTagView struct {
	// Key is the tag key.
	Key *string `json:"key"`
	// Value is the tag value.
	Value *string `json:"value"`
}

// PolicyTemplateListView is a type that runs validations on a projected type.
type PolicyTemplateListView struct {
	// count is the number of policy templates in the list.
	Count *uint
	// etag is an HTTP ETag for the policy template list.
	Etag *string
	// items is the array of policy templates.
	Items PolicyTemplateCollectionView
	// not_modified is a flag used internally that indicates how to encode the HTTP
	// response (i.e. 200 or 304).
	NotModified *string
	// kind is "gov#policy_template_list".
	Kind *string
}

// PolicyTemplateCollectionView is a type that runs validations on a projected
// type.
type PolicyTemplateCollectionView []*PolicyTemplateView

var (
	// PolicyTemplateMap is a map of attribute names in result type PolicyTemplate
	// indexed by view name.
	PolicyTemplateMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"project_id",
			"rs_pt_ver",
			"short_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"required_roles",
			"severity",
			"tenancy",
			"kind",
		},
		"extended": []string{
			"id",
			"name",
			"project_id",
			"rs_pt_ver",
			"short_description",
			"long_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"permissions",
			"required_roles",
			"parameters",
			"severity",
			"tenancy",
			"credentials",
			"kind",
		},
		"source": []string{
			"id",
			"name",
			"href",
			"filename",
			"source",
			"fingerprint",
			"kind",
		},
		"link": []string{
			"id",
			"name",
			"href",
			"fingerprint",
			"updated_at",
			"updated_by",
			"kind",
		},
	}
	// PolicyTemplateListMap is a map of attribute names in result type
	// PolicyTemplateList indexed by view name.
	PolicyTemplateListMap = map[string][]string{
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
	// PolicyTemplateCollectionMap is a map of attribute names in result type
	// PolicyTemplateCollection indexed by view name.
	PolicyTemplateCollectionMap = map[string][]string{
		"default": []string{
			"id",
			"name",
			"project_id",
			"rs_pt_ver",
			"short_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"required_roles",
			"severity",
			"tenancy",
			"kind",
		},
		"extended": []string{
			"id",
			"name",
			"project_id",
			"rs_pt_ver",
			"short_description",
			"long_description",
			"doc_link",
			"info",
			"default_frequency",
			"href",
			"fingerprint",
			"category",
			"created_by",
			"created_at",
			"updated_by",
			"updated_at",
			"permissions",
			"required_roles",
			"parameters",
			"severity",
			"tenancy",
			"credentials",
			"kind",
		},
		"source": []string{
			"id",
			"name",
			"href",
			"filename",
			"source",
			"fingerprint",
			"kind",
		},
		"link": []string{
			"id",
			"name",
			"href",
			"fingerprint",
			"updated_at",
			"updated_by",
			"kind",
		},
	}
)

// ValidatePolicyTemplate runs the validations defined on the viewed result
// type PolicyTemplate.
func ValidatePolicyTemplate(result *PolicyTemplate) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePolicyTemplateView(result.Projected)
	case "extended":
		err = ValidatePolicyTemplateViewExtended(result.Projected)
	case "source":
		err = ValidatePolicyTemplateViewSource(result.Projected)
	case "link":
		err = ValidatePolicyTemplateViewLink(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "extended", "source", "link"})
	}
	return
}

// ValidatePolicyTemplateList runs the validations defined on the viewed result
// type PolicyTemplateList.
func ValidatePolicyTemplateList(result *PolicyTemplateList) (err error) {
	switch result.View {
	case "default", "":
		err = ValidatePolicyTemplateListView(result.Projected)
	case "extended":
		err = ValidatePolicyTemplateListViewExtended(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "extended"})
	}
	return
}

// ValidatePolicyTemplateView runs the validations defined on
// PolicyTemplateView using the "default" view.
func ValidatePolicyTemplateView(result *PolicyTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.DocLink != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.doc_link", *result.DocLink, goa.FormatURI))
	}
	if result.DefaultFrequency != nil {
		if !(*result.DefaultFrequency == "15 minutes" || *result.DefaultFrequency == "hourly" || *result.DefaultFrequency == "daily" || *result.DefaultFrequency == "weekly" || *result.DefaultFrequency == "monthly") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.default_frequency", *result.DefaultFrequency, []interface{}{"15 minutes", "hourly", "daily", "weekly", "monthly"}))
		}
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"))
	}
	if result.CreatedBy != nil {
		if err2 := ValidateUserView(result.CreatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedBy != nil {
		if err2 := ValidateUserView(result.UpdatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	if result.Severity != nil {
		if !(*result.Severity == "low" || *result.Severity == "medium" || *result.Severity == "high" || *result.Severity == "critical") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.severity", *result.Severity, []interface{}{"low", "medium", "high", "critical"}))
		}
	}
	if result.Tenancy != nil {
		if !(*result.Tenancy == "multi" || *result.Tenancy == "single") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.tenancy", *result.Tenancy, []interface{}{"multi", "single"}))
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#policy_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#policy_template"}))
		}
	}
	return
}

// ValidatePolicyTemplateViewExtended runs the validations defined on
// PolicyTemplateView using the "extended" view.
func ValidatePolicyTemplateViewExtended(result *PolicyTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.DocLink != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.doc_link", *result.DocLink, goa.FormatURI))
	}
	if result.DefaultFrequency != nil {
		if !(*result.DefaultFrequency == "15 minutes" || *result.DefaultFrequency == "hourly" || *result.DefaultFrequency == "daily" || *result.DefaultFrequency == "weekly" || *result.DefaultFrequency == "monthly") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.default_frequency", *result.DefaultFrequency, []interface{}{"15 minutes", "hourly", "daily", "weekly", "monthly"}))
		}
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"))
	}
	if result.CreatedBy != nil {
		if err2 := ValidateUserView(result.CreatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.CreatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.created_at", *result.CreatedAt, goa.FormatDateTime))
	}
	if result.UpdatedBy != nil {
		if err2 := ValidateUserView(result.UpdatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	for _, v := range result.Permissions {
		if v != nil {
			if err2 := ValidatePermissionView(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	for _, v := range result.Parameters {
		if v != nil {
			if err2 := ValidateParameterView(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Severity != nil {
		if !(*result.Severity == "low" || *result.Severity == "medium" || *result.Severity == "high" || *result.Severity == "critical") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.severity", *result.Severity, []interface{}{"low", "medium", "high", "critical"}))
		}
	}
	if result.Tenancy != nil {
		if !(*result.Tenancy == "multi" || *result.Tenancy == "single") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.tenancy", *result.Tenancy, []interface{}{"multi", "single"}))
		}
	}
	for _, v := range result.Credentials {
		if v != nil {
			if err2 := ValidateCredentialsView(v); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#policy_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#policy_template"}))
		}
	}
	return
}

// ValidatePolicyTemplateViewSource runs the validations defined on
// PolicyTemplateView using the "source" view.
func ValidatePolicyTemplateViewSource(result *PolicyTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"))
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#policy_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#policy_template"}))
		}
	}
	return
}

// ValidatePolicyTemplateViewLink runs the validations defined on
// PolicyTemplateView using the "link" view.
func ValidatePolicyTemplateViewLink(result *PolicyTemplateView) (err error) {
	if result.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "result"))
	}
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Href == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("href", "result"))
	}
	if result.Fingerprint == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("fingerprint", "result"))
	}
	if result.Kind == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("kind", "result"))
	}
	if result.Href != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.href", *result.Href, "^/api/governance/projects/[0-9]+/policy_templates/[0-9a-f]+$"))
	}
	if result.UpdatedAt != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.updated_at", *result.UpdatedAt, goa.FormatDateTime))
	}
	if result.UpdatedBy != nil {
		if err2 := ValidateUserView(result.UpdatedBy); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if result.Kind != nil {
		if !(*result.Kind == "gov#policy_template") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#policy_template"}))
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

// ValidatePermissionView runs the validations defined on PermissionView.
func ValidatePermissionView(result *PermissionView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Resources == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("resources", "result"))
	}
	if result.Actions == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("actions", "result"))
	}
	return
}

// ValidateParameterView runs the validations defined on ParameterView.
func ValidateParameterView(result *ParameterView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "result"))
	}
	if result.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "result"))
	}
	if result.Index == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("index", "result"))
	}
	if result.Type != nil {
		if !(*result.Type == "string" || *result.Type == "list" || *result.Type == "number") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.type", *result.Type, []interface{}{"string", "list", "number"}))
		}
	}
	if result.AllowedPattern != nil {
		if err2 := ValidateRegexpView(result.AllowedPattern); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateRegexpView runs the validations defined on RegexpView.
func ValidateRegexpView(result *RegexpView) (err error) {
	if result.Pattern == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pattern", "result"))
	}
	return
}

// ValidateCredentialsView runs the validations defined on CredentialsView.
func ValidateCredentialsView(result *CredentialsView) (err error) {
	if result.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "result"))
	}
	if result.Schemes == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("schemes", "result"))
	}
	if result.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "result"))
	}
	for _, e := range result.Tags {
		if e != nil {
			if err2 := ValidateCredentialsTagView(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateCredentialsTagView runs the validations defined on
// CredentialsTagView.
func ValidateCredentialsTagView(result *CredentialsTagView) (err error) {
	if result.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "result"))
	}
	if result.Value == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("value", "result"))
	}
	return
}

// ValidatePolicyTemplateListView runs the validations defined on
// PolicyTemplateListView using the "default" view.
func ValidatePolicyTemplateListView(result *PolicyTemplateListView) (err error) {
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
		if !(*result.Kind == "gov#policy_template_list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#policy_template_list"}))
		}
	}
	if result.Items != nil {
		if err2 := ValidatePolicyTemplateCollectionView(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePolicyTemplateListViewExtended runs the validations defined on
// PolicyTemplateListView using the "extended" view.
func ValidatePolicyTemplateListViewExtended(result *PolicyTemplateListView) (err error) {
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
		if !(*result.Kind == "gov#policy_template_list") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("result.kind", *result.Kind, []interface{}{"gov#policy_template_list"}))
		}
	}
	if result.Items != nil {
		if err2 := ValidatePolicyTemplateCollectionViewExtended(result.Items); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePolicyTemplateCollectionView runs the validations defined on
// PolicyTemplateCollectionView using the "default" view.
func ValidatePolicyTemplateCollectionView(result PolicyTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePolicyTemplateView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePolicyTemplateCollectionViewExtended runs the validations defined on
// PolicyTemplateCollectionView using the "extended" view.
func ValidatePolicyTemplateCollectionViewExtended(result PolicyTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePolicyTemplateViewExtended(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePolicyTemplateCollectionViewSource runs the validations defined on
// PolicyTemplateCollectionView using the "source" view.
func ValidatePolicyTemplateCollectionViewSource(result PolicyTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePolicyTemplateViewSource(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidatePolicyTemplateCollectionViewLink runs the validations defined on
// PolicyTemplateCollectionView using the "link" view.
func ValidatePolicyTemplateCollectionViewLink(result PolicyTemplateCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidatePolicyTemplateViewLink(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
