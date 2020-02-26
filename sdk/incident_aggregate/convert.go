// Code generated by goa v3.0.6, DO NOT EDIT.
//
// IncidentAggregate service type conversion functions
//
// Command:
// $ goa gen github.com/rightscale/governance/front_service/design

package incidentaggregate

import (
	incidentaggregate "github.com/rightscale/governance/incident_aggregate_service/gen/incident_aggregate"
	policyaggregate "github.com/rightscale/governance/policy_aggregate_service/gen/policy_aggregate"
	policytemplate "github.com/rightscale/governance/policy_template_service/gen/policy_template"
	publishedtemplate "github.com/rightscale/governance/published_template_service/gen/published_template"
)

// CreateFromPolicyTemplate initializes t from the fields of v
func (t *PolicyTemplate) CreateFromPolicyTemplate(v *policytemplate.PolicyTemplate) {
	temp := &PolicyTemplate{
		ID:               v.ID,
		Name:             v.Name,
		ProjectID:        v.ProjectID,
		RsPtVer:          v.RsPtVer,
		ShortDescription: v.ShortDescription,
		LongDescription:  v.LongDescription,
		DocLink:          v.DocLink,
		DefaultFrequency: v.DefaultFrequency,
		Href:             v.Href,
		Category:         v.Category,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        v.UpdatedAt,
		Severity:         v.Severity,
		Tenancy:          v.Tenancy,
		Kind:             v.Kind,
	}
	if v.Info != nil {
		temp.Info = make(map[string]string, len(v.Info))
		for key, val := range v.Info {
			tk := key
			tv := val
			temp.Info[tk] = tv
		}
	}
	if v.CreatedBy != nil {
		temp.CreatedBy = transformPolicytemplateUserToUser(v.CreatedBy)
	}
	if v.UpdatedBy != nil {
		temp.UpdatedBy = transformPolicytemplateUserToUser(v.UpdatedBy)
	}
	if v.Permissions != nil {
		temp.Permissions = make(map[string]*Permission, len(v.Permissions))
		for key, val := range v.Permissions {
			tk := key
			temp.Permissions[tk] = transformPolicytemplatePermissionToPermission(val)
		}
	}
	if v.RequiredRoles != nil {
		temp.RequiredRoles = make([]string, len(v.RequiredRoles))
		for i, val := range v.RequiredRoles {
			temp.RequiredRoles[i] = val
		}
	}
	if v.Parameters != nil {
		temp.Parameters = make(map[string]*Parameter, len(v.Parameters))
		for key, val := range v.Parameters {
			tk := key
			temp.Parameters[tk] = transformPolicytemplateParameterToParameter(val)
		}
	}
	if v.Credentials != nil {
		temp.Credentials = make(map[string]*Credentials, len(v.Credentials))
		for key, val := range v.Credentials {
			tk := key
			temp.Credentials[tk] = transformPolicytemplateCredentialsToCredentials(val)
		}
	}
	*t = *temp
}

// CreateFromPolicyAggregateNonCatalog initializes t from the fields of v
func (t *PolicyAggregateNonCatalog) CreateFromPolicyAggregateNonCatalog(v *policyaggregate.PolicyAggregateNonCatalog) {
	temp := &PolicyAggregateNonCatalog{
		Href:                  v.Href,
		IncidentAggregateHref: v.IncidentAggregateHref,
		Count:                 v.Count,
		ActiveCount:           v.ActiveCount,
		ErrorCount:            v.ErrorCount,
		Kind:                  v.Kind,
	}
	if v.RunningProjectIds != nil {
		temp.RunningProjectIds = make([]uint, len(v.RunningProjectIds))
		for i, val := range v.RunningProjectIds {
			temp.RunningProjectIds[i] = val
		}
	}
	if v.Items != nil {
		temp.Items = make([]*PolicyAggregateNonCatalogItem, len(v.Items))
		for i, val := range v.Items {
			temp.Items[i] = transformPolicyaggregatePolicyAggregateNonCatalogItemToPolicyAggregateNonCatalogItem(val)
		}
	}
	*t = *temp
}

// CreateFromPublishedTemplate initializes t from the fields of v
func (t *PublishedTemplate) CreateFromPublishedTemplate(v *publishedtemplate.PublishedTemplate) {
	temp := &PublishedTemplate{
		ID:                        v.ID,
		Name:                      v.Name,
		OrgID:                     v.OrgID,
		ProjectID:                 v.ProjectID,
		PolicyTemplateID:          v.PolicyTemplateID,
		PolicyTemplateURL:         v.PolicyTemplateURL,
		PolicyTemplateFingerprint: v.PolicyTemplateFingerprint,
		RsPtVer:                   v.RsPtVer,
		ShortDescription:          v.ShortDescription,
		LongDescription:           v.LongDescription,
		DocLink:                   v.DocLink,
		DefaultFrequency:          v.DefaultFrequency,
		Href:                      v.Href,
		Category:                  v.Category,
		CreatedAt:                 v.CreatedAt,
		UpdatedAt:                 v.UpdatedAt,
		Severity:                  v.Severity,
		BuiltIn:                   v.BuiltIn,
		Hidden:                    v.Hidden,
		HiddenAt:                  v.HiddenAt,
		Tenancy:                   v.Tenancy,
		Kind:                      v.Kind,
	}
	if v.Info != nil {
		temp.Info = make(map[string]string, len(v.Info))
		for key, val := range v.Info {
			tk := key
			tv := val
			temp.Info[tk] = tv
		}
	}
	if v.CreatedBy != nil {
		temp.CreatedBy = transformPublishedtemplateUserToUser(v.CreatedBy)
	}
	if v.UpdatedBy != nil {
		temp.UpdatedBy = transformPublishedtemplateUserToUser(v.UpdatedBy)
	}
	if v.Permissions != nil {
		temp.Permissions = make(map[string]*Permission, len(v.Permissions))
		for key, val := range v.Permissions {
			tk := key
			temp.Permissions[tk] = transformPublishedtemplatePermissionToPermission(val)
		}
	}
	if v.RequiredRoles != nil {
		temp.RequiredRoles = make([]string, len(v.RequiredRoles))
		for i, val := range v.RequiredRoles {
			temp.RequiredRoles[i] = val
		}
	}
	if v.Parameters != nil {
		temp.Parameters = make(map[string]*Parameter, len(v.Parameters))
		for key, val := range v.Parameters {
			tk := key
			temp.Parameters[tk] = transformPublishedtemplateParameterToParameter(val)
		}
	}
	if v.HiddenBy != nil {
		temp.HiddenBy = transformPublishedtemplateUserToUser(v.HiddenBy)
	}
	if v.Credentials != nil {
		temp.Credentials = make(map[string]*Credentials, len(v.Credentials))
		for key, val := range v.Credentials {
			tk := key
			temp.Credentials[tk] = transformPublishedtemplateCredentialsToCredentials(val)
		}
	}
	*t = *temp
}

// CreateFromIncidentAggregateNonCatalog initializes t from the fields of v
func (t *IncidentAggregateNonCatalog) CreateFromIncidentAggregateNonCatalog(v *incidentaggregate.IncidentAggregateNonCatalog) {
	temp := &IncidentAggregateNonCatalog{
		Href:        &v.Href,
		Count:       v.Count,
		UpdatedAt:   v.UpdatedAt,
		Kind:        v.Kind,
		Etag:        v.Etag,
		NotModified: v.NotModified,
	}
	if v.PolicyAggregate != nil {
		temp.PolicyAggregate = transformIncidentaggregatePolicyAggregateNonCatalogToPolicyAggregateNonCatalog(v.PolicyAggregate)
	}
	if v.IncidentSummary != nil {
		temp.IncidentSummary = transformIncidentaggregateIncidentSummaryToIncidentSummary(v.IncidentSummary)
	}
	if v.ActionSummary != nil {
		temp.ActionSummary = transformIncidentaggregateActionSummaryToActionSummary(v.ActionSummary)
	}
	if v.Items != nil {
		temp.Items = make([]*IncidentAggregateNonCatalogItem, len(v.Items))
		for i, val := range v.Items {
			temp.Items[i] = transformIncidentaggregateIncidentAggregateNonCatalogItemToIncidentAggregateNonCatalogItem(val)
		}
	}
	*t = *temp
}

// transformPolicytemplateUserToUser builds a value of type *User from a value
// of type *policytemplate.User.
func transformPolicytemplateUserToUser(v *policytemplate.User) *User {
	if v == nil {
		return nil
	}
	res := &User{
		ID:    v.ID,
		Email: v.Email,
		Name:  v.Name,
	}

	return res
}

// transformPolicytemplatePermissionToPermission builds a value of type
// *Permission from a value of type *policytemplate.Permission.
func transformPolicytemplatePermissionToPermission(v *policytemplate.Permission) *Permission {
	if v == nil {
		return nil
	}
	res := &Permission{
		Name:  v.Name,
		Label: v.Label,
	}
	if v.Resources != nil {
		res.Resources = make([]string, len(v.Resources))
		for i, val := range v.Resources {
			res.Resources[i] = val
		}
	}
	if v.Actions != nil {
		res.Actions = make([]string, len(v.Actions))
		for i, val := range v.Actions {
			res.Actions[i] = val
		}
	}

	return res
}

// transformPolicytemplateParameterToParameter builds a value of type
// *Parameter from a value of type *policytemplate.Parameter.
func transformPolicytemplateParameterToParameter(v *policytemplate.Parameter) *Parameter {
	if v == nil {
		return nil
	}
	res := &Parameter{
		Name:                  v.Name,
		Type:                  v.Type,
		Label:                 v.Label,
		Index:                 v.Index,
		Category:              v.Category,
		Description:           v.Description,
		Default:               v.Default,
		NoEcho:                v.NoEcho,
		MinLength:             v.MinLength,
		MaxLength:             v.MaxLength,
		MinValue:              v.MinValue,
		MaxValue:              v.MaxValue,
		ConstraintDescription: v.ConstraintDescription,
	}
	if v.AllowedValues != nil {
		res.AllowedValues = make([]interface{}, len(v.AllowedValues))
		for i, val := range v.AllowedValues {
			res.AllowedValues[i] = val
		}
	}
	if v.AllowedPattern != nil {
		res.AllowedPattern = transformPolicytemplateRegexpToRegexp(v.AllowedPattern)
	}

	return res
}

// transformPolicytemplateRegexpToRegexp builds a value of type *Regexp from a
// value of type *policytemplate.Regexp.
func transformPolicytemplateRegexpToRegexp(v *policytemplate.Regexp) *Regexp {
	if v == nil {
		return nil
	}
	res := &Regexp{
		Pattern: v.Pattern,
		Options: v.Options,
	}

	return res
}

// transformPolicytemplateCredentialsToCredentials builds a value of type
// *Credentials from a value of type *policytemplate.Credentials.
func transformPolicytemplateCredentialsToCredentials(v *policytemplate.Credentials) *Credentials {
	if v == nil {
		return nil
	}
	res := &Credentials{
		Name:        v.Name,
		Label:       v.Label,
		Description: v.Description,
	}
	if v.Schemes != nil {
		res.Schemes = make([]string, len(v.Schemes))
		for i, val := range v.Schemes {
			res.Schemes[i] = val
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*CredentialsTag, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformPolicytemplateCredentialsTagToCredentialsTag(val)
		}
	}

	return res
}

// transformPolicytemplateCredentialsTagToCredentialsTag builds a value of type
// *CredentialsTag from a value of type *policytemplate.CredentialsTag.
func transformPolicytemplateCredentialsTagToCredentialsTag(v *policytemplate.CredentialsTag) *CredentialsTag {
	if v == nil {
		return nil
	}
	res := &CredentialsTag{
		Key:   v.Key,
		Value: v.Value,
	}

	return res
}

// transformPolicyaggregatePolicyAggregateNonCatalogItemToPolicyAggregateNonCatalogItem
// builds a value of type *PolicyAggregateNonCatalogItem from a value of type
// *policyaggregate.PolicyAggregateNonCatalogItem.
func transformPolicyaggregatePolicyAggregateNonCatalogItemToPolicyAggregateNonCatalogItem(v *policyaggregate.PolicyAggregateNonCatalogItem) *PolicyAggregateNonCatalogItem {
	if v == nil {
		return nil
	}
	res := &PolicyAggregateNonCatalogItem{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Status:      v.Status,
		Error:       v.Error,
		ErroredAt:   v.ErroredAt,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Severity:    v.Severity,
		Category:    v.Category,
		Frequency:   v.Frequency,
		DryRun:      v.DryRun,
		Kind:        v.Kind,
	}
	if v.Project != nil {
		res.Project = transformPolicyaggregateProjectToProject(v.Project)
	}
	if v.CreatedBy != nil {
		res.CreatedBy = transformPolicyaggregateUserToUser(v.CreatedBy)
	}

	return res
}

// transformPolicyaggregateProjectToProject builds a value of type *Project
// from a value of type *policyaggregate.Project.
func transformPolicyaggregateProjectToProject(v *policyaggregate.Project) *Project {
	if v == nil {
		return nil
	}
	res := &Project{
		ID:      v.ID,
		Name:    v.Name,
		OrgID:   v.OrgID,
		OrgName: v.OrgName,
	}

	return res
}

// transformPolicyaggregateUserToUser builds a value of type *User from a value
// of type *policyaggregate.User.
func transformPolicyaggregateUserToUser(v *policyaggregate.User) *User {
	if v == nil {
		return nil
	}
	res := &User{
		ID:    v.ID,
		Email: v.Email,
		Name:  v.Name,
	}

	return res
}

// transformPublishedtemplateUserToUser builds a value of type *User from a
// value of type *publishedtemplate.User.
func transformPublishedtemplateUserToUser(v *publishedtemplate.User) *User {
	if v == nil {
		return nil
	}
	res := &User{
		ID:    v.ID,
		Email: v.Email,
		Name:  v.Name,
	}

	return res
}

// transformPublishedtemplatePermissionToPermission builds a value of type
// *Permission from a value of type *publishedtemplate.Permission.
func transformPublishedtemplatePermissionToPermission(v *publishedtemplate.Permission) *Permission {
	if v == nil {
		return nil
	}
	res := &Permission{
		Name:  v.Name,
		Label: v.Label,
	}
	if v.Resources != nil {
		res.Resources = make([]string, len(v.Resources))
		for i, val := range v.Resources {
			res.Resources[i] = val
		}
	}
	if v.Actions != nil {
		res.Actions = make([]string, len(v.Actions))
		for i, val := range v.Actions {
			res.Actions[i] = val
		}
	}

	return res
}

// transformPublishedtemplateParameterToParameter builds a value of type
// *Parameter from a value of type *publishedtemplate.Parameter.
func transformPublishedtemplateParameterToParameter(v *publishedtemplate.Parameter) *Parameter {
	if v == nil {
		return nil
	}
	res := &Parameter{
		Name:                  v.Name,
		Type:                  v.Type,
		Label:                 v.Label,
		Index:                 v.Index,
		Category:              v.Category,
		Description:           v.Description,
		Default:               v.Default,
		NoEcho:                v.NoEcho,
		MinLength:             v.MinLength,
		MaxLength:             v.MaxLength,
		MinValue:              v.MinValue,
		MaxValue:              v.MaxValue,
		ConstraintDescription: v.ConstraintDescription,
	}
	if v.AllowedValues != nil {
		res.AllowedValues = make([]interface{}, len(v.AllowedValues))
		for i, val := range v.AllowedValues {
			res.AllowedValues[i] = val
		}
	}
	if v.AllowedPattern != nil {
		res.AllowedPattern = transformPublishedtemplateRegexpToRegexp(v.AllowedPattern)
	}

	return res
}

// transformPublishedtemplateRegexpToRegexp builds a value of type *Regexp from
// a value of type *publishedtemplate.Regexp.
func transformPublishedtemplateRegexpToRegexp(v *publishedtemplate.Regexp) *Regexp {
	if v == nil {
		return nil
	}
	res := &Regexp{
		Pattern: v.Pattern,
		Options: v.Options,
	}

	return res
}

// transformPublishedtemplateCredentialsToCredentials builds a value of type
// *Credentials from a value of type *publishedtemplate.Credentials.
func transformPublishedtemplateCredentialsToCredentials(v *publishedtemplate.Credentials) *Credentials {
	if v == nil {
		return nil
	}
	res := &Credentials{
		Name:        v.Name,
		Label:       v.Label,
		Description: v.Description,
	}
	if v.Schemes != nil {
		res.Schemes = make([]string, len(v.Schemes))
		for i, val := range v.Schemes {
			res.Schemes[i] = val
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*CredentialsTag, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformPublishedtemplateCredentialsTagToCredentialsTag(val)
		}
	}

	return res
}

// transformPublishedtemplateCredentialsTagToCredentialsTag builds a value of
// type *CredentialsTag from a value of type *publishedtemplate.CredentialsTag.
func transformPublishedtemplateCredentialsTagToCredentialsTag(v *publishedtemplate.CredentialsTag) *CredentialsTag {
	if v == nil {
		return nil
	}
	res := &CredentialsTag{
		Key:   v.Key,
		Value: v.Value,
	}

	return res
}

// transformIncidentaggregatePolicyAggregateNonCatalogToPolicyAggregateNonCatalog
// builds a value of type *PolicyAggregateNonCatalog from a value of type
// *incidentaggregate.PolicyAggregateNonCatalog.
func transformIncidentaggregatePolicyAggregateNonCatalogToPolicyAggregateNonCatalog(v *incidentaggregate.PolicyAggregateNonCatalog) *PolicyAggregateNonCatalog {
	if v == nil {
		return nil
	}
	res := &PolicyAggregateNonCatalog{
		Href:                  v.Href,
		IncidentAggregateHref: v.IncidentAggregateHref,
		Count:                 v.Count,
		ActiveCount:           v.ActiveCount,
		ErrorCount:            v.ErrorCount,
		Kind:                  v.Kind,
	}
	if v.RunningProjectIds != nil {
		res.RunningProjectIds = make([]uint, len(v.RunningProjectIds))
		for i, val := range v.RunningProjectIds {
			res.RunningProjectIds[i] = val
		}
	}
	if v.Items != nil {
		res.Items = make([]*PolicyAggregateNonCatalogItem, len(v.Items))
		for i, val := range v.Items {
			res.Items[i] = transformIncidentaggregatePolicyAggregateNonCatalogItemToPolicyAggregateNonCatalogItem(val)
		}
	}

	return res
}

// transformIncidentaggregatePolicyAggregateNonCatalogItemToPolicyAggregateNonCatalogItem
// builds a value of type *PolicyAggregateNonCatalogItem from a value of type
// *incidentaggregate.PolicyAggregateNonCatalogItem.
func transformIncidentaggregatePolicyAggregateNonCatalogItemToPolicyAggregateNonCatalogItem(v *incidentaggregate.PolicyAggregateNonCatalogItem) *PolicyAggregateNonCatalogItem {
	if v == nil {
		return nil
	}
	res := &PolicyAggregateNonCatalogItem{
		ID:          v.ID,
		Name:        v.Name,
		Description: v.Description,
		Status:      v.Status,
		Error:       v.Error,
		ErroredAt:   v.ErroredAt,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
		Severity:    v.Severity,
		Category:    v.Category,
		Frequency:   v.Frequency,
		DryRun:      v.DryRun,
		Kind:        v.Kind,
	}
	if v.Project != nil {
		res.Project = transformIncidentaggregateProjectToProject(v.Project)
	}
	if v.CreatedBy != nil {
		res.CreatedBy = transformIncidentaggregateUserToUser(v.CreatedBy)
	}

	return res
}

// transformIncidentaggregateProjectToProject builds a value of type *Project
// from a value of type *incidentaggregate.Project.
func transformIncidentaggregateProjectToProject(v *incidentaggregate.Project) *Project {
	if v == nil {
		return nil
	}
	res := &Project{
		ID:      v.ID,
		Name:    v.Name,
		OrgID:   v.OrgID,
		OrgName: v.OrgName,
	}

	return res
}

// transformIncidentaggregateUserToUser builds a value of type *User from a
// value of type *incidentaggregate.User.
func transformIncidentaggregateUserToUser(v *incidentaggregate.User) *User {
	if v == nil {
		return nil
	}
	res := &User{
		ID:    v.ID,
		Email: v.Email,
		Name:  v.Name,
	}

	return res
}

// transformIncidentaggregateIncidentSummaryToIncidentSummary builds a value of
// type *IncidentSummary from a value of type
// *incidentaggregate.IncidentSummary.
func transformIncidentaggregateIncidentSummaryToIncidentSummary(v *incidentaggregate.IncidentSummary) *IncidentSummary {
	if v == nil {
		return nil
	}
	res := &IncidentSummary{
		IncidentCount:      v.IncidentCount,
		ViolationDataCount: v.ViolationDataCount,
		ResolvedCount:      v.ResolvedCount,
		TriggeredCount:     v.TriggeredCount,
	}

	return res
}

// transformIncidentaggregateActionSummaryToActionSummary builds a value of
// type *ActionSummary from a value of type *incidentaggregate.ActionSummary.
func transformIncidentaggregateActionSummaryToActionSummary(v *incidentaggregate.ActionSummary) *ActionSummary {
	if v == nil {
		return nil
	}
	res := &ActionSummary{
		PendingCount: v.PendingCount,
		FailedCount:  v.FailedCount,
	}

	return res
}

// transformIncidentaggregateIncidentAggregateNonCatalogItemToIncidentAggregateNonCatalogItem
// builds a value of type *IncidentAggregateNonCatalogItem from a value of type
// *incidentaggregate.IncidentAggregateNonCatalogItem.
func transformIncidentaggregateIncidentAggregateNonCatalogItemToIncidentAggregateNonCatalogItem(v *incidentaggregate.IncidentAggregateNonCatalogItem) *IncidentAggregateNonCatalogItem {
	if v == nil {
		return nil
	}
	res := &IncidentAggregateNonCatalogItem{
		ID:                 v.ID,
		CreatedAt:          v.CreatedAt,
		UpdatedAt:          v.UpdatedAt,
		ResolvedAt:         v.ResolvedAt,
		ResolutionMessage:  v.ResolutionMessage,
		State:              v.State,
		ViolationDataCount: v.ViolationDataCount,
		ActionFailed:       v.ActionFailed,
		ActionPending:      v.ActionPending,
		Severity:           v.Severity,
		Category:           v.Category,
		DryRun:             v.DryRun,
		Kind:               v.Kind,
	}
	if v.Project != nil {
		res.Project = transformIncidentaggregateProjectToProject(v.Project)
	}
	if v.AppliedPolicy != nil {
		res.AppliedPolicy = transformIncidentaggregateAppliedPolicyToAppliedPolicy(v.AppliedPolicy)
	}
	if v.ResolvedBy != nil {
		res.ResolvedBy = transformIncidentaggregateUserToUser(v.ResolvedBy)
	}

	return res
}

// transformIncidentaggregateAppliedPolicyToAppliedPolicy builds a value of
// type *AppliedPolicy from a value of type *incidentaggregate.AppliedPolicy.
func transformIncidentaggregateAppliedPolicyToAppliedPolicy(v *incidentaggregate.AppliedPolicy) *AppliedPolicy {
	if v == nil {
		return nil
	}
	res := &AppliedPolicy{
		ID:                  v.ID,
		PolicyAggregateID:   v.PolicyAggregateID,
		IncidentAggregateID: v.IncidentAggregateID,
		Name:                v.Name,
		Href:                v.Href,
		Description:         v.Description,
		DocLink:             v.DocLink,
		CreatedAt:           v.CreatedAt,
		UpdatedAt:           v.UpdatedAt,
		Severity:            v.Severity,
		Category:            v.Category,
		Frequency:           v.Frequency,
		DryRun:              v.DryRun,
		SkipApprovals:       v.SkipApprovals,
		Status:              v.Status,
		Error:               v.Error,
		ErroredAt:           v.ErroredAt,
		Scope:               v.Scope,
		Kind:                v.Kind,
	}
	if v.Project != nil {
		res.Project = transformIncidentaggregateProjectToProject(v.Project)
	}
	if v.PolicyTemplate != nil {
		res.PolicyTemplate = transformIncidentaggregatePolicyTemplateToPolicyTemplate(v.PolicyTemplate)
	}
	if v.PublishedTemplate != nil {
		res.PublishedTemplate = transformIncidentaggregatePublishedTemplateToPublishedTemplate(v.PublishedTemplate)
	}
	if v.Info != nil {
		res.Info = make(map[string]string, len(v.Info))
		for key, val := range v.Info {
			tk := key
			tv := val
			res.Info[tk] = tv
		}
	}
	if v.CreatedBy != nil {
		res.CreatedBy = transformIncidentaggregateUserToUser(v.CreatedBy)
	}
	if v.Options != nil {
		res.Options = make([]*ConfigurationOption, len(v.Options))
		for i, val := range v.Options {
			res.Options[i] = transformIncidentaggregateConfigurationOptionToConfigurationOption(val)
		}
	}
	if v.Credentials != nil {
		res.Credentials = make(map[string]string, len(v.Credentials))
		for key, val := range v.Credentials {
			tk := key
			tv := val
			res.Credentials[tk] = tv
		}
	}

	return res
}

// transformIncidentaggregatePolicyTemplateToPolicyTemplate builds a value of
// type *PolicyTemplate from a value of type *incidentaggregate.PolicyTemplate.
func transformIncidentaggregatePolicyTemplateToPolicyTemplate(v *incidentaggregate.PolicyTemplate) *PolicyTemplate {
	if v == nil {
		return nil
	}
	res := &PolicyTemplate{
		ID:               v.ID,
		Name:             v.Name,
		ProjectID:        v.ProjectID,
		RsPtVer:          v.RsPtVer,
		ShortDescription: v.ShortDescription,
		LongDescription:  v.LongDescription,
		DocLink:          v.DocLink,
		DefaultFrequency: v.DefaultFrequency,
		Href:             v.Href,
		Category:         v.Category,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        v.UpdatedAt,
		Severity:         v.Severity,
		Tenancy:          v.Tenancy,
		Kind:             v.Kind,
	}
	if v.Info != nil {
		res.Info = make(map[string]string, len(v.Info))
		for key, val := range v.Info {
			tk := key
			tv := val
			res.Info[tk] = tv
		}
	}
	if v.CreatedBy != nil {
		res.CreatedBy = transformIncidentaggregateUserToUser(v.CreatedBy)
	}
	if v.UpdatedBy != nil {
		res.UpdatedBy = transformIncidentaggregateUserToUser(v.UpdatedBy)
	}
	if v.Permissions != nil {
		res.Permissions = make(map[string]*Permission, len(v.Permissions))
		for key, val := range v.Permissions {
			tk := key
			res.Permissions[tk] = transformIncidentaggregatePermissionToPermission(val)
		}
	}
	if v.RequiredRoles != nil {
		res.RequiredRoles = make([]string, len(v.RequiredRoles))
		for i, val := range v.RequiredRoles {
			res.RequiredRoles[i] = val
		}
	}
	if v.Parameters != nil {
		res.Parameters = make(map[string]*Parameter, len(v.Parameters))
		for key, val := range v.Parameters {
			tk := key
			res.Parameters[tk] = transformIncidentaggregateParameterToParameter(val)
		}
	}
	if v.Credentials != nil {
		res.Credentials = make(map[string]*Credentials, len(v.Credentials))
		for key, val := range v.Credentials {
			tk := key
			res.Credentials[tk] = transformIncidentaggregateCredentialsToCredentials(val)
		}
	}

	return res
}

// transformIncidentaggregatePermissionToPermission builds a value of type
// *Permission from a value of type *incidentaggregate.Permission.
func transformIncidentaggregatePermissionToPermission(v *incidentaggregate.Permission) *Permission {
	if v == nil {
		return nil
	}
	res := &Permission{
		Name:  v.Name,
		Label: v.Label,
	}
	if v.Resources != nil {
		res.Resources = make([]string, len(v.Resources))
		for i, val := range v.Resources {
			res.Resources[i] = val
		}
	}
	if v.Actions != nil {
		res.Actions = make([]string, len(v.Actions))
		for i, val := range v.Actions {
			res.Actions[i] = val
		}
	}

	return res
}

// transformIncidentaggregateParameterToParameter builds a value of type
// *Parameter from a value of type *incidentaggregate.Parameter.
func transformIncidentaggregateParameterToParameter(v *incidentaggregate.Parameter) *Parameter {
	if v == nil {
		return nil
	}
	res := &Parameter{
		Name:                  v.Name,
		Type:                  v.Type,
		Label:                 v.Label,
		Index:                 v.Index,
		Category:              v.Category,
		Description:           v.Description,
		Default:               v.Default,
		NoEcho:                v.NoEcho,
		MinLength:             v.MinLength,
		MaxLength:             v.MaxLength,
		MinValue:              v.MinValue,
		MaxValue:              v.MaxValue,
		ConstraintDescription: v.ConstraintDescription,
	}
	if v.AllowedValues != nil {
		res.AllowedValues = make([]interface{}, len(v.AllowedValues))
		for i, val := range v.AllowedValues {
			res.AllowedValues[i] = val
		}
	}
	if v.AllowedPattern != nil {
		res.AllowedPattern = transformIncidentaggregateRegexpToRegexp(v.AllowedPattern)
	}

	return res
}

// transformIncidentaggregateRegexpToRegexp builds a value of type *Regexp from
// a value of type *incidentaggregate.Regexp.
func transformIncidentaggregateRegexpToRegexp(v *incidentaggregate.Regexp) *Regexp {
	if v == nil {
		return nil
	}
	res := &Regexp{
		Pattern: v.Pattern,
		Options: v.Options,
	}

	return res
}

// transformIncidentaggregateCredentialsToCredentials builds a value of type
// *Credentials from a value of type *incidentaggregate.Credentials.
func transformIncidentaggregateCredentialsToCredentials(v *incidentaggregate.Credentials) *Credentials {
	if v == nil {
		return nil
	}
	res := &Credentials{
		Name:        v.Name,
		Label:       v.Label,
		Description: v.Description,
	}
	if v.Schemes != nil {
		res.Schemes = make([]string, len(v.Schemes))
		for i, val := range v.Schemes {
			res.Schemes[i] = val
		}
	}
	if v.Tags != nil {
		res.Tags = make([]*CredentialsTag, len(v.Tags))
		for i, val := range v.Tags {
			res.Tags[i] = transformIncidentaggregateCredentialsTagToCredentialsTag(val)
		}
	}

	return res
}

// transformIncidentaggregateCredentialsTagToCredentialsTag builds a value of
// type *CredentialsTag from a value of type *incidentaggregate.CredentialsTag.
func transformIncidentaggregateCredentialsTagToCredentialsTag(v *incidentaggregate.CredentialsTag) *CredentialsTag {
	if v == nil {
		return nil
	}
	res := &CredentialsTag{
		Key:   v.Key,
		Value: v.Value,
	}

	return res
}

// transformIncidentaggregatePublishedTemplateToPublishedTemplate builds a
// value of type *PublishedTemplate from a value of type
// *incidentaggregate.PublishedTemplate.
func transformIncidentaggregatePublishedTemplateToPublishedTemplate(v *incidentaggregate.PublishedTemplate) *PublishedTemplate {
	if v == nil {
		return nil
	}
	res := &PublishedTemplate{
		ID:                        v.ID,
		Name:                      v.Name,
		OrgID:                     v.OrgID,
		ProjectID:                 v.ProjectID,
		PolicyTemplateID:          v.PolicyTemplateID,
		PolicyTemplateURL:         v.PolicyTemplateURL,
		PolicyTemplateFingerprint: v.PolicyTemplateFingerprint,
		RsPtVer:                   v.RsPtVer,
		ShortDescription:          v.ShortDescription,
		LongDescription:           v.LongDescription,
		DocLink:                   v.DocLink,
		DefaultFrequency:          v.DefaultFrequency,
		Href:                      v.Href,
		Category:                  v.Category,
		CreatedAt:                 v.CreatedAt,
		UpdatedAt:                 v.UpdatedAt,
		Severity:                  v.Severity,
		BuiltIn:                   v.BuiltIn,
		Hidden:                    v.Hidden,
		HiddenAt:                  v.HiddenAt,
		Tenancy:                   v.Tenancy,
		Kind:                      v.Kind,
	}
	if v.Info != nil {
		res.Info = make(map[string]string, len(v.Info))
		for key, val := range v.Info {
			tk := key
			tv := val
			res.Info[tk] = tv
		}
	}
	if v.CreatedBy != nil {
		res.CreatedBy = transformIncidentaggregateUserToUser(v.CreatedBy)
	}
	if v.UpdatedBy != nil {
		res.UpdatedBy = transformIncidentaggregateUserToUser(v.UpdatedBy)
	}
	if v.Permissions != nil {
		res.Permissions = make(map[string]*Permission, len(v.Permissions))
		for key, val := range v.Permissions {
			tk := key
			res.Permissions[tk] = transformIncidentaggregatePermissionToPermission(val)
		}
	}
	if v.RequiredRoles != nil {
		res.RequiredRoles = make([]string, len(v.RequiredRoles))
		for i, val := range v.RequiredRoles {
			res.RequiredRoles[i] = val
		}
	}
	if v.Parameters != nil {
		res.Parameters = make(map[string]*Parameter, len(v.Parameters))
		for key, val := range v.Parameters {
			tk := key
			res.Parameters[tk] = transformIncidentaggregateParameterToParameter(val)
		}
	}
	if v.HiddenBy != nil {
		res.HiddenBy = transformIncidentaggregateUserToUser(v.HiddenBy)
	}
	if v.Credentials != nil {
		res.Credentials = make(map[string]*Credentials, len(v.Credentials))
		for key, val := range v.Credentials {
			tk := key
			res.Credentials[tk] = transformIncidentaggregateCredentialsToCredentials(val)
		}
	}

	return res
}

// transformIncidentaggregateConfigurationOptionToConfigurationOption builds a
// value of type *ConfigurationOption from a value of type
// *incidentaggregate.ConfigurationOption.
func transformIncidentaggregateConfigurationOptionToConfigurationOption(v *incidentaggregate.ConfigurationOption) *ConfigurationOption {
	if v == nil {
		return nil
	}
	res := &ConfigurationOption{
		Name:  v.Name,
		Label: v.Label,
		Type:  v.Type,
		Value: v.Value,
	}

	return res
}
